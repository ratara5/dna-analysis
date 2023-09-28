package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ratara5/dna-analysis/routes"
)

func main() {
	// Crea una instancia de Gin
	r := gin.Default()

	routes.DnaRouter(r)

	// Ejecuta la aplicación en el puerto 8080
	r.Run(":8080")
}