# 对于go的垃圾回收调整

更多的是在于对其内存大小的限制以及对于其内存占用的限制，go的gc存在着活跃内存以及新生内存，新生内存可能是被临时抽调使用，所以在运行的时候，通常会对cpu的最大使用率进行限制，特别是gc的内存占用，否则非常容易导致oom的出现。

也就是在这种情况，并发回收所产生的stw有时候会导致整个程序崩溃，所以通常的gogc参数不能设置为100或者-1[解开限制]。

但是并不是简单的限制其gogc的大小或者一昧的提高gogc的大小能提升效率，不根据cpu来进行调整往往会对整个运行效率产生更坏的效果。下面是go官网对于gogc参数的描述：

~~~makefile
关于 GOGC 的附加说明
GOGC 部分声称， 将GOGC 翻倍会使堆内存开销翻倍，并将 GC CPU 成本减半。要了解原因，让我们在数学上对其进行分解。

首先，堆目标为总堆大小设置一个目标。然而，这个目标主要影响新的堆内存，因为活动堆是应用程序的基础。

目标堆内存 = 活动堆 +（活动堆 + GC 根）* GOGC / 100

总堆内存 = 活动堆 + 新堆内存

⇒

新堆内存 =（活动堆 + GC 根）* GOGC / 100

由此我们可以看到，将 GOGC 翻倍也会使应用程序在每个周期分配的新堆内存量翻倍，这会捕获堆内存开销。请注意，实时堆 + GC 根是 GC 需要扫描的内存量的近似值。

接下来，我们来看看 GC CPU 开销。总成本可以分解为每个周期的成本乘以一段时间 T 内的 GC 频率。

总 GC CPU 成本 = (每个周期的 GC CPU 成本) * (GC 频率) * T

每个周期的 GC CPU 成本可以从 GC 模型中得出：

每个周期的 GC CPU 成本 =（活动堆 + GC 根）*（每字节成本）+ 固定成本

请注意，此处忽略扫描阶段成本，因为标记和扫描成本占主导地位。

稳态由恒定的分配率和恒定的每字节成本定义，因此在稳态下，我们可以从这个新的堆内存中推导出 GC 频率：

GC 频率 = (分配率) / (新堆内存) = (分配率) / ((活堆 + GC 根) * GOGC / 100)

把这些放在一起，我们得到了总成本的完整方程：

总 GC CPU 成本 = (分配率) / ((活堆 + GC 根) * GOGC / 100) * ((活堆 + GC 根) * (每字节成本) + 固定成本) * T

对于足够大的堆（代表大多数情况），GC 循环的边际成本支配着固定成本。这可以显着简化 GC CPU 总成本公式。

总 GC CPU 成本 = (分配率) / (GOGC / 100) * (每字节成本) * T

从这个简化的公式中，我们可以看到，如果我们将 GOGC 翻倍，我们将总 GC CPU 成本减半。（请注意，本指南中的可视化确实模拟了固定成本，因此当 GOGC 翻倍时，它们报告的 GC CPU 开销不会完全减半。）此外，GC CPU 成本在很大程度上取决于分配率和扫描内存的每字节成本。有关如何具体降低这些成本的更多信息，请参阅 优化指南。

注意：活动堆的大小和 GC 实际需要扫描的内存量之间存在差异：相同大小的活动堆但具有不同的结构会导致不同的 CPU 成本，但相同的内存成本，导致不同的权衡。这就是为什么堆的结构是稳态定义的一部分。堆目标可以说应该只包括可扫描的活动堆，作为 GC 需要扫描的内存的更接近的近似值，但是当可扫描的活动堆数量非常少但活动堆很大时，这会导致退化行为。
~~~
