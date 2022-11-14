package main

import (
	"database/sql"
	"log"
)

// AlbumByID retrieves the specified album.
// 你期望重复执行同一条SQL时，可以使用ansql.Stmt 提前准备好SQL语句，然后按需执行。
// 以下示例创建一个准备好的语句，从数据库中选择一个特定的专辑。DB.Prepare 返回sql.Stmt表示给定 SQL 文本的准备语句。您可以将 SQL 语句的参数传递给Stmt.Exec、Stmt.QueryRow或Stmt.Query以运行该语句。
func AlbumByID(id int) (Album, error) {
	// Define a prepared statement. You'd typically define the statement
	// elsewhere and save it for use in functions such as this one.
	stmt, err := db.Prepare("SELECT * FROM album WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	var album Album

	// Execute the prepared statement, passing in an id value for the
	// parameter whose placeholder is ?
	err = stmt.QueryRow(id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case of no rows returned.
		}
		return album, err
	}
	return album, nil
}
