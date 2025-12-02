package presentation_technique

import (
	application_technique "bjj-tracker/src/modules/technique/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TechniqueHandler struct {
	CreateTechniqueUC  *application_technique.CreateTechniqueUseCase
	GetTechniqueByIDUC *application_technique.GetTechniqueByIDUseCase
	GetAllTechniquesUC *application_technique.GetAllTechniquesUseCase
	UpdateTechniqueUC  *application_technique.UpdateTechniqueUseCase
}

func NewTechniqueHandler(createTechniqueUC *application_technique.CreateTechniqueUseCase, updateTechniqueUC *application_technique.UpdateTechniqueUseCase, getTechniqueByIDUC *application_technique.GetTechniqueByIDUseCase, getAllTechniquesUC *application_technique.GetAllTechniquesUseCase) *TechniqueHandler {
	return &TechniqueHandler{
		CreateTechniqueUC:  createTechniqueUC,
		UpdateTechniqueUC:  updateTechniqueUC,
		GetTechniqueByIDUC: getTechniqueByIDUC,
		GetAllTechniquesUC: getAllTechniquesUC,
	}
}

func (h *TechniqueHandler) CreateTechnique(c *gin.Context) {
	var req application_technique.CreateTechniqueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	technique, err := h.CreateTechniqueUC.Execute(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"technique": technique})
}

func (h *TechniqueHandler) UpdateTechnique(c *gin.Context) {
	id := c.Param("id")
	var req application_technique.UpdateTechniqueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	technique, err := h.UpdateTechniqueUC.Execute(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"technique": technique})
	return
}

func (h *TechniqueHandler) GetTechniqueByID(c *gin.Context) {
	id := c.Param("id")
	technique, err := h.GetTechniqueByIDUC.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"technique": technique})
}

func (h *TechniqueHandler) GetAllTechniques(c *gin.Context) {
	techniques, err := h.GetAllTechniquesUC.Execute()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"techniques": techniques})
}
