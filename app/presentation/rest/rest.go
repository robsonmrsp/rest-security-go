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

	movieRepo := repository.NewRepo(db)

	routers.Movies(engine, *movieRepo, mid)

	engine.Run()
}
