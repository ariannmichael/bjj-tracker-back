package presentation_technique

import (
	"jiu-tracker/middleware"

	"github.com/gin-gonic/gin"
)

func TechniqueRoutes(r *gin.RouterGroup, handler *TechniqueHandler) {
	r.POST("/technique", middleware.RequireAuth, handler.CreateTechnique)
	r.PUT("/technique/:id", middleware.RequireAuth, handler.UpdateTechnique)
	r.GET("/technique/:id", middleware.RequireAuth, handler.GetTechniqueByID)
	r.GET("/techniques", middleware.RequireAuth, handler.GetAllTechniques)
	r.GET("/techniques/list", middleware.RequireAuth, handler.GetTechniquesList)
}
