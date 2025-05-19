package presentation_user

import (
	application_user "bjj-tracker/src/modules/user/application"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	CreateUserUC *application_user.CreateUserUseCase
	LoginUserUC  *application_user.LoginUserUseCase
}

func NewUserHandler(createUserUC *application_user.CreateUserUseCase, loginUserUC *application_user.LoginUserUseCase) *UserHandler {
	return &UserHandler{
		CreateUserUC: createUserUC,
		LoginUserUC:  loginUserUC,
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	req.Password = string(hash)
	user, err := h.CreateUserUC.Execute(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := h.LoginUserUC.Execute(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", *tokenString, 3600*24, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{})
}

func (h *UserHandler) Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
	})
}
