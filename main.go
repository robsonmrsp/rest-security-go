package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robsonmrsp/rest-security-go/controllers"
)

func main() {
	fmt.Println("Hello!")
	router := gin.Default()

	personGroup := router.Group("rs/crud/person")
	personGroup.GET("/:id", controllers.GetPerson)
	personGroup.POST("/", controllers.PostPerson)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
