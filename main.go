package main

import (
	"fmt"

	mysql_api "example.com/m/v2/mysql"
	"example.com/m/v2/router"
)

func main() {
	err:=mysql_api.InitDB()//数据库初始化
	if err != nil {
		fmt.Println("failed")
	}else {
	fmt.Println("success")	
	}
	 router.InitRoute()//初始化路由
	 
}