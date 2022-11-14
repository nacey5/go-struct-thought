package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB = Connect()

func Connect() *sql.DB {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "a1160124552",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return db
}

func Conncet2() *sql.DB {
	// Specify connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "a1160124552",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "jazzrecords",
	}

	// Get a driver-specific connector.
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Get a database handle.
	db = sql.OpenDB(connector)
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}
