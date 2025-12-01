package main

import (
	"bjj-tracker/config"
	application_technique "bjj-tracker/src/modules/technique/application"
	presentation_technique "bjj-tracker/src/modules/technique/presentation"
	application_training "bjj-tracker/src/modules/training/application"
	presentation_training "bjj-tracker/src/modules/training/presentation"
	application_user "bjj-tracker/src/modules/user/application"
	presentation_user "bjj-tracker/src/modules/user/presentation"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func initUserCases() (*application_user.CreateUserUseCase, *application_user.LoginUserUseCase, *application_user.GetUserByIDUseCase, *application_user.GetAllUsersUseCase) {
	createUserUC := application_user.NewCreateUserUseCase()
	loginUserUC := application_user.NewLoginUserUseCase()
	getUserByIDUC := application_user.NewGetUserByIDUseCase()
	getAllUsersUC := application_user.NewGetAllUsersUseCase()
	return createUserUC, loginUserUC, getUserByIDUC, getAllUsersUC
}

func initUserHandler(createUserUC *application_user.CreateUserUseCase, loginUserUC *application_user.LoginUserUseCase, getUserByIDUC *application_user.GetUserByIDUseCase, getAllUsersUC *application_user.GetAllUsersUseCase) *presentation_user.UserHandler {
	return &presentation_user.UserHandler{
		CreateUserUC:  createUserUC,
		LoginUserUC:   loginUserUC,
		GetUserByIDUC: getUserByIDUC,
		GetAllUsersUC: getAllUsersUC,
	}
}

func initTrainingCases() *application_training.CreateTrainingUseCase {
	createTrainingUC := application_training.NewCreateTrainingUseCase()
	return createTrainingUC
}

func initTrainingHandler(createTrainingUC *application_training.CreateTrainingUseCase) *presentation_training.TrainingHandler {
	return &presentation_training.TrainingHandler{
		CreateTrainingUC: createTrainingUC,
	}
}

func initTechniqueCases() *application_technique.CreateTechniqueUseCase {
	createTechniqueUC := application_technique.NewCreateTechniqueUseCase()
	return createTechniqueUC
}

func initTechniqueHandler(createTechniqueUC *application_technique.CreateTechniqueUseCase) *presentation_technique.TechniqueHandler {
	return &presentation_technique.TechniqueHandler{
		CreateTechniqueUC: createTechniqueUC,
	}
}

func main() {
	router := gin.Default()

	// User routes
	createUserUC, loginUserUC, getUserByIDUC, getAllUsersUC := initUserCases()
	userHandler := initUserHandler(createUserUC, loginUserUC, getUserByIDUC, getAllUsersUC)
	presentation_user.UserRoutes(router.Group("/api"), userHandler)

	// Training routes
	createTrainingUC := initTrainingCases()
	trainingHandler := initTrainingHandler(createTrainingUC)
	presentation_training.TrainingRoutes(router.Group("/api"), trainingHandler)

	// Technique routes
	createTechniqueUC := initTechniqueCases()
	techniqueHandler := initTechniqueHandler(createTechniqueUC)
	presentation_technique.TechniqueRoutes(router.Group("/api"), techniqueHandler)

	router.Run()
}
