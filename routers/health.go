package routers

import (
	"fapi/controllers"

	"github.com/gin-gonic/gin"
)

func HealthRoute(router *gin.Engine) {
	status:= new(controllers.HealthRepo)
router.GET("/health", status.Health)
}