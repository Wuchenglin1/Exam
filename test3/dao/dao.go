package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

const user = "root:root@/test1"

func InitMySql() {
	db, err := sql.Open("mysql", user)
	if err != nil {
		panic(err)
	}
	dB = db
}
