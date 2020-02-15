package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConn, err := sql.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
}
