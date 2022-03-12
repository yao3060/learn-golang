package main

import (
	"booking/controller"
	"booking/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	r := gin.Default()
	// r.Use(gin.Recovery(), middlleware.Logger())

	r.GET("/posts", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})

	r.POST("/posts", func(c *gin.Context) {
		c.JSON(201, videoController.Save(c))
	})

	r.Run(":8000")
}
