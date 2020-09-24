package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
	"github.com/robsonmrsp/rest-security-go/app/domain/services"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/parser"
)

type movieController struct {
	service services.MovieService
}

// MovieController ...
type MovieController interface {
	GetMovie() func(c *gin.Context)
	PostMovie() func(c *gin.Context)
	GetPageMovie() func(c *gin.Context)
	PutMovie() func(c *gin.Context)
	DeleteMovie() func(c *gin.Context)
}

// NewMovieController ...
func NewMovieController(movieService services.MovieService) MovieController {
	return &movieController{movieService}
}

func (controller *movieController) GetMovie() func(c *gin.Context) {
	return func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		person, _ := controller.service.GetMovie(id)
		context.JSON(200, person)
	}
}

func (controller *movieController) GetPageMovie() func(context *gin.Context) {
	return func(context *gin.Context) {

		p := parser.Create(context)
		movie, _ := controller.service.GetPageMovie(p)

		context.JSON(200, movie)
	}
}

func (controller *movieController) PostMovie() func(context *gin.Context) {
	return func(context *gin.Context) {
		mo := &entities.Movie{}
		err := context.BindJSON(&mo)
		if err != nil {
			context.JSON(500, err)
		}
		mo, err = controller.service.SaveMovie(mo)
		if err != nil {
			context.JSON(500, err)

		}
		context.JSON(201, mo)
	}
}

func (controller *movieController) PutMovie() func(context *gin.Context) {
	return func(context *gin.Context) {

	}
}

func (controller *movieController) DeleteMovie() func(context *gin.Context) {
	return func(context *gin.Context) {

	}
}
