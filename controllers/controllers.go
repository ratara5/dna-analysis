package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ratara5/dna-analysis/analize"
	"net/http"
)

type RequestBody struct {
	Atributo []string `json:"dna"`
}

func VerifyIsMutant(c *gin.Context) {
	var requestBody RequestBody

	// Intenta analizar el JSON de la solicitud en la estructura RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verifica la condición en el atributo

	if analize.IsMutant(requestBody.Atributo) { //TODO: Mover IsMutant de main.go
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
	}

}