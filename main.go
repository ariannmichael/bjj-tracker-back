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

	userHandler := presentation_user.NewUserHandler(&application_user.CreateUserUseCase{})
	apiGroup := router.Group("/api")
	presentation_user.UserRoutes(apiGroup, userHandler)

	router.Run()
}
