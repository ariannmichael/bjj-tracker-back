package presentation_technique

import (
	"bjj-tracker/middleware"

	"github.com/gin-gonic/gin"
)

func TechniqueRoutes(r *gin.RouterGroup, handler *TechniqueHandler) {
	r.POST("/technique", middleware.RequireAuth, handler.CreateTechnique)
	r.GET("/technique/:id", middleware.RequireAuth, handler.GetTechniqueByID)
	r.GET("/techniques", middleware.RequireAuth, handler.GetAllTechniques)
}
