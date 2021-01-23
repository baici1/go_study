package mysql_api

import (
	"fmt"

	"example.com/m/v2/dao"
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
func 	QueryRowDemo(name string) (user dao.User,err error) {
	sqlstr:="select id,name,password from user where name=?"
	var u dao.User
	err=db.Get(&u,sqlstr,name)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return u,err
	}
	//fmt.Println(u)
	return u,nil
}
//增加数据库

func InsertRowDemo(name string,password string) error {
	sqlstr:="insert into user(name, password) values (?,?)"
	_,err:=db.Exec(sqlstr,name,password)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return nil
}