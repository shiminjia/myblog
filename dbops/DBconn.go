package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

var (
	DB *sql.DB
	err error
)

func InitDB(){
	DB, err = sql.Open("mysql", "root:11111111@tcp(localhost:3306)/myblog_dev?charset=utf8")
	if err != nil {
		log.Panicln(err.Error())
	}
	err = DB.Ping()
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Println("数据库 myblog_dev 连接成功")
}

