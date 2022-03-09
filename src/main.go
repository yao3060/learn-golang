package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.yaoyingying.com",
			"wechat": "yao306015195",
			"author": "Yao Yingying",
		})
	})
	r.Run(":8000")
}
