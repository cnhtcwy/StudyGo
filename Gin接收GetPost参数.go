package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func login(c *gin.Context) {
	var user User
	fmt.Println(c.PostForm("username"))
	fmt.Println(c.PostForm("password"))
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"password": user.Password,
	})
}
func main() {
	r := gin.Default()
	r.GET("/test/:name", func(c *gin.Context) {
		username := c.Query("username")
		name := c.Param("name")          //路由地址为：/user/:name/:pass，获取参数
		reqPara := c.Request.URL.Query() //获取Get的所有参数
		for k, v := range reqPara {
			fmt.Println(k, v)
		}
		//fmt.Println(username)
		c.JSON(200, gin.H{
			"msg": name + ":hello world:" + username,
		})
	})
	r.POST("/test", func(c *gin.Context) {
		name := c.PostForm("name")
		price := c.DefaultPostForm("price", "100")
		reqPost := c.Request.PostForm //获取 Post 所有参数
		for k, v := range reqPost {
			println(k, &v)
		}
		c.JSON(200, gin.H{
			"name":  name,
			"price": price,
		})
	})
	r.POST("/login", login)
	r.Run()
}
