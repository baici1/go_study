package mysql_api

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

//连接数据库
func InitDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db,err=sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return 
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return 
}

//查询数据库

//增加数据库