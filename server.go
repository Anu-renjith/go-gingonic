package main

import (
	"github.com/Anu-renjith/go-gingonic/controller"
	"github.com/Anu-renjith/go-gingonic/service"
	"github.com/gin-gonic/gin"
)

var (
	VideoService    service.VideoService    = service.New()
	VideoController controller.VideoController = controller.New(VideoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8082")
}
