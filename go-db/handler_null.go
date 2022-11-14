package main

import (
	"database/sql"
	"log"
)

// 该包提供了几种特殊类型，当列的值可能为空时database/sql，您可以将它们用作函数的参数。Scan每个都包括一个Valid报告该值是否为非空的字段，以及一个保存该值的字段（如果是）。
//
// 以下示例中的代码查询客户名称。如果 name 值为 null，则代码将替换另一个值以在应用程序中使用。
func handlerNull(id int) {
	var s sql.NullString
	err := db.QueryRow("SELECT name FROM customer WHERE id = ?", id).Scan(&s)
	if err != nil {
		log.Fatal(err)
	}
	// Find customer name, using placeholder if not present.
	name := "Valued Customer"
	if s.Valid {
		name = s.String
	}
	println(name)
}
