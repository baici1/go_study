package api

import (
	"net/http"

	mysql_api "example.com/m/v2/mysql"
	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) {
	name := c.Query("name")
	u,err:=mysql_api.QueryRowDemo(name)
	if err != nil {
		c.JSON(404, gin.H{
			"message":"未找到此用户",
		})
		return
	}
		c.JSON(http.StatusOK, gin.H{
		"message":"OK",
		"name": u.Name,
		"password":u.Password,
		"id":u.Id,
		})
	
}

func Register(c *gin.Context)  {
	name := c.Query("name")
	password := c.Query("password")
	_,err:=mysql_api.QueryRowDemo(name)
	if err == nil {
		c.JSON(500, gin.H{
			"message":"用户已存在",
		})
		return
	}
	err=mysql_api.InsertRowDemo(name,password)
	if err != nil {
		c.JSON(500,gin.H{
			"message":"注册出错",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"OK",
	})
}