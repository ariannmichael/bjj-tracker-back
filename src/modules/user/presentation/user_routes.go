package presentation_user

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.RouterGroup, handler *UserHandler) {
	r.POST("/user", handler.CreateUser)
	// r.GET("/user/:id", handler.GetUserByID)
	// r.PUT("/user/:id", handler.UpdateUser)
	// r.DELETE("/user/:id", handler.DeleteUser)
	// r.GET("/users", handler.GetAllUsers)
	// r.POST("/user/login", handler.LoginUser)
	// r.POST("/user/logout", handler.LogoutUser)
	// r.POST("/user/refresh-token", handler.RefreshToken)
	// r.POST("/user/forgot-password", handler.ForgotPassword)
	// r.POST("/user/reset-password", handler.ResetPassword)
}
