package presentation_user

import (
	"bjj-tracker/config"
	application_user "bjj-tracker/src/modules/user/application"
	infrastructure_user "bjj-tracker/src/modules/user/infrastructure"

	"net/http"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	CreateUserUC *application_user.CreateUserUseCase
}

func NewUserHandler(createUserUC *application_user.CreateUserUseCase) *UserHandler {
	db := config.ConnectToDB()
	repo := &infrastructure_user.UserRepositoryImpl{DB: db}
	useCase := &application_user.CreateUserUseCase{Repo: repo}
	return &UserHandler{
		CreateUserUC: useCase,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.CreateUserUC.Execute(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
