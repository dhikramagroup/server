package main

import (
	"io"
	"net/http"
	"os"

	"github.com/dhikramagroup/gin-server/controller"
	"github.com/dhikramagroup/gin-server/midlewares"
	"github.com/dhikramagroup/gin-server/services"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService     services.VideosService      = services.New()
	videosController controller.VideosController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()
	server := gin.New()

	server.Static("/css", ".template/css")
	server.LoadHTMLGlob("templates/*.html")
	server.Use(gin.Recovery(), midlewares.Logger(),
		midlewares.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(c *gin.Context) {
			c.JSON(200, videosController.FindAll())
		})

		apiRoutes.POST("/videos", func(c *gin.Context) {
			err := videosController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Video input valid"})
			}
		})
	}

	viewRoute := server.Group("/view")
	{
		viewRoute.GET("/videos", videosController.ShowAll)
	}

	server.Run(":3000")
}
