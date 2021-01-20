package router

import "github.com/gin-gonic/gin"

func InitRoute() {
	router:=gin.Default()

	v:=router.Group("admin")
	{
		v.GET("/login", )
		v.GET("/register",)
	}
	router.Run()
}