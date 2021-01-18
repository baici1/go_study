package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var DB *sql.DB

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

 func initDB() error {
	var err error
	dsn:="root:123456@tcp(127.0.0.1:3306)/sql_test"
	//使用mysql驱动
	DB,err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// defer db.Close()
	//尝试与数据库连接
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

//CURD
// 查询单条数据示例
func queryResult() {
	querySQL := "SELECT `id`,`name`,`age` FROM `test_user` WHERE `id`=?"
	// 执行查询语句
	row := DB.QueryRow(querySQL, 1)
    // 定义 user 为结构体 User 类型
	var user User
	// 此处获取结果的顺序一定要和 SELECT 语句取出的顺序一样
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		fmt.Printf("query result failed, err: %v\n", err)
		return
	}
	fmt.Printf("query result: id=%d name=%s age=%d\n", user.Id, user.Name, user.Age)
}


func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}else {
		fmt.Println("success")
	}
	 queryResult()
}