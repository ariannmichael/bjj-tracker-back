package presentation_training

import (
	"bjj-tracker/middleware"

	"github.com/gin-gonic/gin"
)

func TrainingRoutes(r *gin.RouterGroup, handler *TrainingHandler) {
	r.POST("/training", middleware.RequireAuth, handler.CreateTraining)
	// r.GET("/training/:id", middleware.RequireAuth, handler.GetTrainingByID)
	// r.GET("/trainings", middleware.RequireAuth, handler.GetAllTrainings)
}
