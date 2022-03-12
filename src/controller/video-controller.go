package controller

import (
	"booking/entity"
	"booking/service"

	"github.com/gin-gonic/gin"
)

type VideoController struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return VideoController{
		service: service,
	}
}

func (c *VideoController) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *VideoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
