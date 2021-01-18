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
	DB,err = sql.Open("mysql", dsn)
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
func queryRowDemo()  {
	sqlStr:="select id,name,age from test_user where id=?"
	var u User
	err := DB.QueryRow(sqlStr,1).Scan(&u.Id,&u.Name,&u.Age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return	
	}
	fmt.Printf("id:%d name:%s age:%d\n",u.Id,u.Name,u.Age)
}


//查询多条数据
func queryMultiRowDemo(id int)  {
	sqlStr:="select id, name, age from test_user where id >= ?"
	rows,err:= DB.Query(sqlStr,id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return	
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	for rows.Next(){
		var u User
		err:=rows.Scan(&u.Id,&u.Name,&u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n",u.Id,u.Name,u.Age)
	}
}


//插入、更新和删除操作都使用Exec方法。
//插入数据
func insertRowDemo(name string,age int)  {
	sqlStr:="insert into test_user(name, age) values (?,?)"
	ret,err:=DB.Exec(sqlStr,name,age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theid ,err:=ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theid)
}

//更新数据
func updateRowDemo(age int,id int)  {
	sqlStr:="update test_user set age=? where id = ?"
	ret,err:= DB.Exec(sqlStr,age,id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n,err:= ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}
//删除数据
func deleteRowDemo(id int)  {
	sqlStr:="delete from test_user where id = ?"
	ret,err:=DB.Exec(sqlStr,id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n,err:= ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}


func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}else {
		fmt.Println("success")
	}
	queryRowDemo()
	queryMultiRowDemo(1)
	insertRowDemo("yayyy",21)
	updateRowDemo(18,3)
	deleteRowDemo(1)
}