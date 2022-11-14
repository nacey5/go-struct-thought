package main

import (
	"database/sql"
	"fmt"
)

// 查询一行
// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// QueryRow最多检索单个数据库行，例如当您要通过唯一 ID 查找数据时。如果查询返回多行，则该 Scan方法将丢弃除第一行之外的所有行。
//
// QueryRowContext像QueryRow但有一个context.Context论点。有关更多信息，请参阅取消正在进行的操作。
//
// 以下示例使用查询来确定是否有足够的库存来支持购买。true如果有足够的 SQL 语句返回，false 如果没有。通过指针Row.Scan将布尔返回值复制到变量中。enough
func canPurchase(id int, quantity int) (bool, error) {
	var enough bool
	// Query for a value based on a single row.
	if err := db.QueryRow("SELECT (quantity >= ?) from album where id = ?",
		quantity, id).Scan(&enough); err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("canPurchase %d: unknown album", id)
		}
		return false, fmt.Errorf("canPurchase %d: %v", id)
	}
	return enough, nil
}
