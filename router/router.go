package router

import (
	"example.com/m/v2/api"
	"example.com/m/v2/middlewares"
	"github.com/gin-gonic/gin"
)


func InitRoute() {
	router:=gin.Default()
	router.Use(middlewares.Cors())
	v:=router.Group("admin")
	{
		v.POST("/login", api.Login)
		v.POST("/register",api.Register)
	}
	router.Run(":8090")
	
}