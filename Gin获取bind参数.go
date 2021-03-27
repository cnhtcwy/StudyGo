package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Age      int       `form:"age" binging:"required,gt=10"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/test", testing)
	r.POST("/test", testing)
	r.Run()
}
func testing(c *gin.Context) {
	var person Person
	//根据content-type来binding操作
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(200, "person bind error%v", err)
	}
}
