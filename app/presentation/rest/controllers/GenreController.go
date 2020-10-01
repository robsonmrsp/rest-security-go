package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
	"github.com/robsonmrsp/rest-security-go/app/domain/services"
)

type genreController struct {
	service services.GenreService
}

// GenreController ...
type GenreController interface {
	GetGenre() func(c *gin.Context)
	PostGenre() func(c *gin.Context)
	GetPageGenre() func(c *gin.Context)
	PutGenre() func(c *gin.Context)
	DeleteGenre() func(c *gin.Context)
}

// NewGenreController ...
func NewGenreController(genreService services.GenreService) GenreController {
	return &genreController{genreService}
}

func (controller *genreController) GetGenre() func(c *gin.Context) {
	return func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		person, _ := controller.service.GetGenre(id)
		context.JSON(200, person)
	}
}

func (controller *genreController) GetPageGenre() func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("name")

		context.JSON(200, "devolvendo o id "+id)
	}
}

func (controller *genreController) PostGenre() func(context *gin.Context) {
	return func(context *gin.Context) {
		mo := &entities.Genre{}
		err := context.BindJSON(&mo)
		if err != nil {
			context.JSON(500, err)
		} else {
			mo, err = controller.service.SaveGenre(mo)
			if err != nil {
				context.JSON(500, err)
			} else {
				context.JSON(201, mo)
			}
		}
	}
}

func (controller *genreController) PutGenre() func(context *gin.Context) {
	return func(context *gin.Context) {
		mo := &entities.Genre{}
		err := context.BindJSON(&mo)
		if err != nil {
			context.JSON(500, err)
		}
		mo, err = controller.service.UpdateGenre(mo)
		if err != nil {
			context.JSON(500, err)
		}
		context.JSON(201, mo)
	}
}

func (controller *genreController) DeleteGenre() func(context *gin.Context) {
	return func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		mo, err = controller.service.DeleteGenre(id)
		context.JSON(204)

	}
}
