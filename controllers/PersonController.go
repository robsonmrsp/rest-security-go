package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/robsonmrsp/rest-security-go/services"
	"strconv"
)

// GetPerson ...
func GetPerson(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	person, _ := services.GetPerson(id)
	context.JSON(200, person)
}

// GetPagePerson ...
func GetPagePerson(context *gin.Context) {
	id := context.Param("name")
	context.JSON(200, "devolvendo o id "+id)
}

// PostPerson ...
func PostPerson(context *gin.Context) {

}

// PutPerson ...
func PutPerson(context *gin.Context) {
}

// DeletePerson ..
func DeletePerson(context *gin.Context) {
}
