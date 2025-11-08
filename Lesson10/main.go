package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法1：使用map
		// data := map[string]interface{}{
		// 	"name":    "jackie",
		// 	"message": "hello world",
		// 	"age":     20,
		// }

		//与上述方式定义本质上一样，gin框架做了封装
		data := gin.H{
			"name":    "jackie",
			"message": "hello world",
			"age":     20,
		}
		c.JSON(http.StatusOK, data)
	})
	//方法2：结构体（字段小写为不可导出，注意这个坑），灵活使用tag来对结构体字段做定制化操作
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}
	r.GET("another_json", func(c *gin.Context) {
		data := msg{
			"jackie",
			"Hello golang",
			20,
		}
		c.JSON(http.StatusOK, data) //json的序列化
	})
	r.Run(":8080")
}
