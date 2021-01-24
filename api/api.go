package api

import (
	"net/http"

	mysql_api "example.com/m/v2/mysql"
	jwt "example.com/m/v2/tool"
	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) {
	name := c.Query("name")
	// fmt.Println(name)
	u,err:=mysql_api.QueryRowDemo(name)
	//fmt.Println(u)
	if err != nil {
		c.JSON(404, gin.H{
			"message":"未找到此用户",
			"code":"404",
		})
		return
	}
	token,_:=jwt.GenToken(u.Id,u.Name,u.Password)
	//fmt.Println(token)
		c.JSON(http.StatusOK, gin.H{
		"message":"OK",
		"data":gin.H{
			"name": u.Name,
			"password":u.Password,
			"id":u.Id,
			"token":token,
		},
		"code":"200",
		})
	
}

func Register(c *gin.Context)  {
	name := c.Query("name")
	password := c.Query("password")
	u,err:=mysql_api.QueryRowDemo(name)
	if err == nil {
		c.JSON(500, gin.H{
			"message":"用户已存在",
			"code":"500",
		})
		return
	}
	err=mysql_api.InsertRowDemo(name,password)
	if err != nil {
		c.JSON(500,gin.H{
			"message":"注册出错",
			"code":"500",
		})
		return
	}
	u,err=mysql_api.QueryRowDemo(name)
token,_:=jwt.GenToken(u.Id,u.Name,u.Password)
	//fmt.Println(token)
		c.JSON(http.StatusOK, gin.H{
		"message":"OK",
		"data":gin.H{
			"name": u.Name,
			"password":u.Password,
			"id":u.Id,
			"token":token,
		},
		"code":"200",
		})
}