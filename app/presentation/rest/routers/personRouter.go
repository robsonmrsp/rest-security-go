package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/robsonmrsp/rest-security-go/app/presentation/rest/controllers"
)

// People ... Create a router group to rs/crud/persons and relative routes
func People(engine *gin.Engine, midlewares ...gin.HandlerFunc) {
	personGroup := engine.Group("rs/crud/person")

	personGroup.GET("/:id", controllers.GetPerson)
	personGroup.GET("/", controllers.GetPagePerson)
	personGroup.PUT("/:id", controllers.PutPerson)
	personGroup.DELETE("/:id", controllers.DeletePerson)
}
