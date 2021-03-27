package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//1.返回json
//c.JSON(200,gin.H{"msg":"OK"})
//c.JSON(200,结构体)
//2.返回模版
//router.LoadHTMLGlob("templates/**/*")
//router.GET("/test/index", func(c *gin.Context){
//	c.HTML(http.StatusOK, "test/index.tmpl", gin.H{
//		"msg": "test",
//	})
//})
//模版index.tmpl
/**
{{define "test/index.tmpl"}}
<html>

    <head>
    </head>

    <body>

        test...

        {{.}}
        -----
        {{.msg}}

    </body>

</html>

{{end}}
*/
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
func uploads(c *gin.Context) {
	//多文件
	//获取解析后表单
	form, _ := c.MultipartForm()
	//这里是多文件上传 在之前单文件upload上传的基础上加 [] 变成upload[] 类似文件数组的意思
	files := form.File["upload[]"]
	//files := form.File["file"]
	//循环存文件到服务器本地
	for _, file := range files {
		fmt.Println(file.Filename)
		//dst:=fmt.Sprintf("D:/GO/src/github.com/teacherwyang@163.com/StudyGo/uploads/%s",file.Filename)
		dst := fmt.Sprintf("./uploads/%s", file.Filename)
		fmt.Println(dst)
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	}
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
	r.POST("/uploads", uploads)
	r.Run()
}
