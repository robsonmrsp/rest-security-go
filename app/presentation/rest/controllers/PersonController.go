package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robsonmrsp/rest-security-go/app/domain/entities"
	"github.com/robsonmrsp/rest-security-go/app/domain/services"
)

// GetPerson ...
func GetPerson(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))

	// aki vai ser iniciado a transação
	person, _ := services.GetPerson(id)
	// aqui termina! talvez usando o "defer"

	context.JSON(200, person)
}

// GetPagePerson ...
func GetPagePerson(context *gin.Context) {
	id := context.Param("name")

	context.JSON(200, "devolvendo o id "+id)
}

// PostPerson ...
func PostPerson(context *gin.Context) {
	p := &entities.Person{}
	err := context.BindJSON(&p)
	if err != nil {
		context.JSON(500, err)
	}
	p, err = services.SavePerson(p)
	if err != nil {
		context.JSON(500, err)
		
	}
	context.JSON(201, p)
}

// PutPerson ...
func PutPerson(context *gin.Context) {
}

// DeletePerson ..
func DeletePerson(context *gin.Context) {
}
