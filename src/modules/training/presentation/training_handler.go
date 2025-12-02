package presentation_training

import (
	application_training "bjj-tracker/src/modules/training/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainingHandler struct {
	CreateTrainingUC  *application_training.CreateTrainingUseCase
	GetTrainingByIDUC *application_training.GetTrainingByIDUseCase
	GetAllTrainingsUC *application_training.GetAllTrainingsUseCase
	UpdateTrainingUC  *application_training.UpdateTrainingUseCase
	DeleteTrainingUC  *application_training.DeleteTrainingUseCase
}

func NewTrainingHandler(createTrainingUC *application_training.CreateTrainingUseCase, getTrainingByIDUC *application_training.GetTrainingByIDUseCase, getAllTrainingsUC *application_training.GetAllTrainingsUseCase, updateTrainingUC *application_training.UpdateTrainingUseCase, deleteTrainingUC *application_training.DeleteTrainingUseCase) *TrainingHandler {
	return &TrainingHandler{
		CreateTrainingUC:  createTrainingUC,
		GetTrainingByIDUC: getTrainingByIDUC,
		GetAllTrainingsUC: getAllTrainingsUC,
		UpdateTrainingUC:  updateTrainingUC,
		DeleteTrainingUC:  deleteTrainingUC,
	}
}

func (h *TrainingHandler) CreateTraining(c *gin.Context) {
	var req application_training.CreateTrainingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	training, err := h.CreateTrainingUC.Execute(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"training": training})
}

func (h *TrainingHandler) GetTrainingByID(c *gin.Context) {
	id := c.Param("id")
	training, err := h.GetTrainingByIDUC.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"training": training})
}

func (h *TrainingHandler) GetAllTrainings(c *gin.Context) {
	trainings, err := h.GetAllTrainingsUC.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"trainings": trainings})
}

func (h *TrainingHandler) UpdateTraining(c *gin.Context) {
	id := c.Param("id")
	var req application_training.UpdateTrainingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	training, err := h.UpdateTrainingUC.Execute(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"training": training})
	return
}

func (h *TrainingHandler) DeleteTraining(c *gin.Context) {
	id := c.Param("id")
	err := h.DeleteTrainingUC.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Training deleted successfully"})
	return
}
