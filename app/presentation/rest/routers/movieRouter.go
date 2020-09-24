package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/robsonmrsp/rest-security-go/app/data/repository"
	"github.com/robsonmrsp/rest-security-go/app/domain/services"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/controllers"
)

// Movies ...
func Movies(engine *gin.Engine, repo repository.MovieRepository, middlewares gin.HandlerFunc) {

	service := services.NewMovieService(repo)

	controller := controllers.NewMovieController(service)

	movieGroup := engine.Group("rs/crud/movies")
	movieGroup.Use(middlewares)

	movieGroup.GET("/:id", controller.GetMovie())
	movieGroup.GET("/", controller.GetPageMovie())
	movieGroup.PUT("/:id", controller.PutMovie())
	movieGroup.DELETE("/:id", controller.DeleteMovie())
}
