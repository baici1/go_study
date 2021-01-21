package router

import (
	"example.com/m/v2/api"
	"github.com/gin-gonic/gin"
)


func InitRoute() {
	router:=gin.Default()

	v:=router.Group("admin")
	{
		v.GET("/login", api.Login)
		v.GET("/register",api.Register)
	}
	router.Run()
}