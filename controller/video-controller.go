package controller

import (
	"github.com/Anu-renjith/go-gingonic/entity"
	"github.com/Anu-renjith/go-gingonic/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}
