package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitSqliteConnection() {
	d, err := sql.Open("sqlite3", "./e-commerce.db")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected To Sqlite DB")
	}

	db = d
}

func GetDB() *sql.DB {
	return db
}
