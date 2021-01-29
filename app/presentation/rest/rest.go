package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/robsonmrsp/rest-security-go/app/data/repository"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/middlewares"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/routers"
)

// Start calls all nedd to startup the application. Database access and Rest API
func Start() {
	engine := gin.Default()
	db := repository.GetDb()
	mid := middlewares.DBTransactionMidleware(db)

	movieRepo := repository.NewMovieRepository(db)
	genreRepo := repository.NewGenreRepository(db)

	routers.Movies(engine, movieRepo, mid)
	routers.Genres(engine, genreRepo, mid)

	engine.Run(":8888")
}
