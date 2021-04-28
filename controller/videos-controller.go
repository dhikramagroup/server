package controller

import (
	"net/http"

	"github.com/dhikramagroup/gin-server/entity"
	"github.com/dhikramagroup/gin-server/services"
	"github.com/dhikramagroup/gin-server/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideosController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	services services.VideosService
}

var validate *validator.Validate

func New(services services.VideosService) VideosController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidatorTitleCool)
	return &controller{
		services: services,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.services.FindAll()
}
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.services.Save(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.services.FindAll()
	data := gin.H{
		"title": "Video page",
		"video": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
