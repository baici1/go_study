package router

import (
	handler_jwt "example.com/m/v2/HandlerFunc"
	"example.com/m/v2/api"
	"github.com/gin-gonic/gin"
)


func InitRoute() {
	router:=gin.Default()

	v:=router.Group("admin")
	{
		v.GET("/login",handler_jwt.JWTAuthMiddleware, api.Login)
		v.POST("/register",api.Register)
	}
	router.Run()
}