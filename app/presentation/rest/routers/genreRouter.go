package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/robsonmrsp/rest-security-go/app/data/repository"
	"github.com/robsonmrsp/rest-security-go/app/domain/services"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/controllers"
)

// Genres ...
func Genres(engine *gin.Engine, repo repository.GenreRepository, middlewares gin.HandlerFunc) {

	service := services.NewGenreService(repo)

	controller := controllers.NewGenreController(service)

	genreGroup := engine.Group("rs/crud/genres")
	genreGroup.Use(middlewares)

	genreGroup.GET("/:id", controller.GetGenre())
	genreGroup.GET("/", controller.GetPageGenre())
	genreGroup.PUT("/:id", controller.PutGenre())
	genreGroup.DELETE("/:id", controller.DeleteGenre())
}
