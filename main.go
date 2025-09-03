package main

import (
	"bjj-tracker/config"
	application_user "bjj-tracker/src/modules/user/application"
	presentation_user "bjj-tracker/src/modules/user/presentation"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	router := gin.Default()

	createUserUC := application_user.NewCreateUserUseCase()
	loginUserUC := application_user.NewLoginUserUseCase()
	getUserByIDUC := application_user.NewGetUserByIDUseCase()
	getAllUsersUC := application_user.NewGetAllUsersUseCase()
	userHandler := presentation_user.NewUserHandler(createUserUC, loginUserUC, getUserByIDUC, getAllUsersUC)
	apiGroup := router.Group("/api")
	presentation_user.UserRoutes(apiGroup, userHandler)

	router.Run()
}
