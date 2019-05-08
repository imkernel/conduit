package api

import (
	"github.com/gin-gonic/gin"
)

func createRoutes(r *gin.Engine) {
	adminGroup := r.Group("admin")
	adminGroupV1 := adminGroup.Group("v1")
	adminGroupV1.POST("/api/application", createApplicationHandler)
}
