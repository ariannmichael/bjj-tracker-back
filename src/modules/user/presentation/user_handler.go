package presentation_user

import (
	application_user "bjj-tracker/src/modules/user/application"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	CreateUserUC *application_user.CreateUserUseCase
}

func NewUserHandler(createUserUC *application_user.CreateUserUseCase) *UserHandler {
	return &UserHandler{
		CreateUserUC: createUserUC,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Name       string `json:"name" binding:"required"`
		Username   string `json:"username" binding:"required"`
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required,min=6"`
		BeltColor  string `json:"belt_color" binding:"required"`
		BeltStripe int    `json:"belt_stripe" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.CreateUserUC.Execute(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"user": user})
}
