package presentation_training

import (
	application_training "bjj-tracker/src/modules/training/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainingHandler struct {
	CreateTrainingUC *application_training.CreateTrainingUseCase
}

func NewTrainingHandler(createTrainingUC *application_training.CreateTrainingUseCase) *TrainingHandler {
	return &TrainingHandler{
		CreateTrainingUC: createTrainingUC,
	}
}

func (h *TrainingHandler) CreateTraining(c *gin.Context) {
	var req application_training.CreateTrainingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	training, err := h.CreateTrainingUC.Execute(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"training": training})
}
