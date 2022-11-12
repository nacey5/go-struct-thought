# go-struct之泛型

在go语言中，泛型的定义必须指定其泛型所能够承受的范围，比如官方给予的例子
~~~go
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
// Initialize a map for the integer values
ints := map[string]int64{
"first":  34,
"second": 12,
}

// Initialize a map for the float values
floats := map[string]float64{
"first":  35.98,
"second": 26.99,
}

fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
SumIntsOrFloats(ints),
SumIntsOrFloats(floats))
~~~

也就是当你正在使用泛型时，所指定的类型必须限制范围，如`map[string]int64` or `map[string]float64`
在泛型中必可以变为 `map[K comparable,V int64 | float64]`,当当前的KV没有指定范围的时候
编译器的编译并不会通过。  
----------------------
因此,我去寻找为何如此实现，且go的底层的泛型是如何实现的，我看到了一篇较为不错的文章:
[Go泛型是怎么实现的？](https://cn-sec.com/archives/487357.html)


