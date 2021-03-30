package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

// 定义一个中间件
func MyFMT() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		fmt.Printf("Before: %s\n", host)
		c.Next()
		fmt.Println("Next: ...")
	}
}
func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 注册一个全局中间件
	r.Use(StatCost())

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	//为某个路由单独注册
	// 给/test2路由单独注册中间件（可注册多个）
	r.GET("/test2", MyFMT(), func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	//为路由组注册中间件
	//方法1
	//shopGroup := r.Group("/shop", StatCost())
	//{
	//	shopGroup.GET("/index", func(c *gin.Context) {...})
	//	...
	//}
	//方法2
	//shopGroup := r.Group("/shop")
	//shopGroup.Use(StatCost())
	//{
	//	shopGroup.GET("/index", func(c *gin.Context) {...})
	//	...
	//}
	r.Run()
}
