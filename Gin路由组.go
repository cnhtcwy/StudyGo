package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"logout": "success",
	})
}
func login1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"login": "success",
	})
}
func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/logout", logout)
		userGroup.GET("/login", login1)
		userGroup.POST("/login", login1)

	}
	//shopGroup := r.Group("/shop")
	//{
	//	shopGroup.GET("/index", func(c *gin.Context) {...})
	//	shopGroup.GET("/cart", func(c *gin.Context) {...})
	//	shopGroup.POST("/checkout", func(c *gin.Context) {...})
	//}

	//shopGroup := r.Group("/shop")
	//{
	//	shopGroup.GET("/index", func(c *gin.Context) {...})
	//	shopGroup.GET("/cart", func(c *gin.Context) {...})
	//	shopGroup.POST("/checkout", func(c *gin.Context) {...})
	//	// 嵌套路由组
	//	xx := shopGroup.Group("xx")
	//	xx.GET("/oo", func(c *gin.Context) {...})
	//}
	r.Run()
}
