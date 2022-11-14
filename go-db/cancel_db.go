package main

import (
	"context"
	"log"
	"time"
)

// 超时后取消数据库操作
// 您可以使用 aContext设置将取消操作的超时或截止日期。要导出Context具有超时或截止日期的 a，请调用 context.WithTimeout或 context.WithDeadline。
//
// 以下超时示例中的代码派生 aContext并将其传递给sql.DB QueryContext 方法。
func QueryWithTimeout(ctx context.Context) {
	// Create a Context with a timeout.
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Pass the timeout Context with a query.
	rows, err := db.QueryContext(queryCtx, "SELECT * FROM album")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Handle returned rows.
}
