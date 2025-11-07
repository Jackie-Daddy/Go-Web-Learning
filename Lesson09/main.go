package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 静态文件：html页面上用到的样式文件 .css js文件 图片
func main() {
	r := gin.Default()
	r.Static("/xxx", "./statics")
	//gin框架中给模版添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// r.LoadHTMLFiles("templates/index.tmpl", "templates/users/index.tmpl") //模版解析
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ //模版渲染
			"title": "www.google.com",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ //模版渲染
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	r.Run(":8080") //启动server
}
