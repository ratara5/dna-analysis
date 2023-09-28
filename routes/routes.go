package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ratara5/dna-analysis/controllers"
)

func DnaRouter(router *gin.Engine) {
	routes := router.Group("")
	routes.POST("/mutant/", controllers.VerifyIsMutant)

}
