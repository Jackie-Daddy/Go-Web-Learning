package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Lesson17 路由与路由组

func main() {
	r := gin.Default()
	//访问/index的GET请求会走着一条处理逻辑
	//路由
	// r.HEAD()  HEAD请求
	//GET请求一般用于获取信息
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	//POST请求一般用于创建信息
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	//DELETE请求一般用于删除信息
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	//PUT请求一般用于更新信息
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	//Any: 请求方法的大集合/大杂烩
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case "POST":
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
			})
			//...
		}
	})
	//NoRoute
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "jackie",
		})
	})
	//视频的首页和详情页
	// r.GET("/video/index", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "/video/index",
	// 	})
	// })
	// r.GET("/video/index", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "/video/index",
	// 	})
	// })
	// r.GET("/video/index", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "/video/index",
	// 	})
	// })

	//路由组 多用于区分不同的业务线或API版本
	//把公用的前缀提取出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/xx",
			})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/oo",
			})
		})
	}

	//商城的首页和详情页
	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "/shop/index",
		})
	})
	r.GET("/shop/xx", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "/shop/xx",
		})
	})
	r.GET("/shop/oo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "/shop/oo",
		})
	})
	r.Run(":8080")
}
