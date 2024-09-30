package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDbClient() *sql.DB {
	client, err := sql.Open("mysql", "root:thecodecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
