package controller

import (
	"github.com/Anu-renjith/go-gingonic/entity"
	"github.com/Anu-renjith/go-gingonic/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

// FindAll implements VideoController.
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// Save implements VideoController.
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	if err := ctx.BindJSON(&video); err != nil {
		return entity.Video{} // Handle error properly, maybe log it or return a proper response
	}
	return c.service.Save(video)
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}
