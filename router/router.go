package router

import "github.com/gin-gonic/gin"

func SetRouter(r *gin.Engine) {
	r.GET("/apis/data", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "hello world!",
		})
	})
}
