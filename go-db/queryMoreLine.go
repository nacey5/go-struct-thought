package main

import (
	"fmt"
)

//查询多行

// Album 对应mysql中的表
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// 查询多行
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

// Query您可以使用or查询多行QueryContext，它返回Rows表示查询结果的 a。您的代码使用 . 遍历返回的行Rows.Next。每次迭代调用Scan将列值复制到变量中。
//
// QueryContext像Query但有一个context.Context论点。有关更多信息，请参阅取消正在进行的操作。
//
// 以下示例执行查询以返回指定艺术家的专辑。专辑以sql.Rows. 该代码用于 Rows.Scan将列值复制到由指针表示的变量中。
func albumsByArtist2(artist string) ([]Album, error) {
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var albums []Album

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist,
			&alb.Price); err != nil {
			return albums, err
		}
		albums = append(albums, alb)
	}
	if err = rows.Err(); err != nil {
		return albums, err
	}
	return albums, nil
}
