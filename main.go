package main

import (
	"jiu-tracker/config"
	application_technique "jiu-tracker/src/modules/technique/application"
	infrastructure_technique "jiu-tracker/src/modules/technique/infrastructure"
	presentation_technique "jiu-tracker/src/modules/technique/presentation"
	application_training "jiu-tracker/src/modules/training/application"
	infrastructure_training "jiu-tracker/src/modules/training/infrastructure"
	presentation_training "jiu-tracker/src/modules/training/presentation"
	application_user "jiu-tracker/src/modules/user/application"
	presentation_user "jiu-tracker/src/modules/user/presentation"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func initUserCases() (*application_user.CreateUserUseCase, *application_user.UpdateUserByIDUseCase, *application_user.LoginUserUseCase, *application_user.GetUserByIDUseCase, *application_user.GetAllUsersUseCase) {
	createUserUC := application_user.NewCreateUserUseCase()
	updateUserByIDUC := application_user.NewUpdateUserByIDUseCase()
	loginUserUC := application_user.NewLoginUserUseCase()
	getUserByIDUC := application_user.NewGetUserByIDUseCase()
	getAllUsersUC := application_user.NewGetAllUsersUseCase()
	return createUserUC, updateUserByIDUC, loginUserUC, getUserByIDUC, getAllUsersUC
}

func initUserHandler(createUserUC *application_user.CreateUserUseCase, updateUserByIDUC *application_user.UpdateUserByIDUseCase, loginUserUC *application_user.LoginUserUseCase, getUserByIDUC *application_user.GetUserByIDUseCase, getAllUsersUC *application_user.GetAllUsersUseCase) *presentation_user.UserHandler {
	return &presentation_user.UserHandler{
		CreateUserUC:     createUserUC,
		UpdateUserByIDUC: updateUserByIDUC,
		LoginUserUC:      loginUserUC,
		GetUserByIDUC:    getUserByIDUC,
		GetAllUsersUC:    getAllUsersUC,
	}
}

func initTrainingHandler() *presentation_training.TrainingHandler {
	db := config.DB
	trainingRepo := &infrastructure_training.TrainingRepositoryImpl{DB: db}
	techniqueRepo := &infrastructure_technique.TechniqueRepositoryImpl{DB: db}
	techniqueService := application_technique.NewTechniqueService(techniqueRepo)

	createTrainingUC := application_training.NewCreateTrainingUseCaseWithDeps(trainingRepo, techniqueService)
	getTrainingByIDUC := application_training.NewGetTrainingByIDUseCase(trainingRepo)
	getAllTrainingsUC := application_training.NewGetAllTrainingsUseCase(trainingRepo)
	updateTrainingUC := application_training.NewUpdateTrainingUseCase(trainingRepo, techniqueService)
	deleteTrainingUC := application_training.NewDeleteTrainingUseCase(trainingRepo)

	return presentation_training.NewTrainingHandler(
		createTrainingUC,
		getTrainingByIDUC,
		getAllTrainingsUC,
		updateTrainingUC,
		deleteTrainingUC,
	)
}

func initTechniqueHandler() *presentation_technique.TechniqueHandler {
	db := config.DB
	techniqueRepo := &infrastructure_technique.TechniqueRepositoryImpl{DB: db}

	createTechniqueUC := application_technique.NewCreateTechniqueUseCaseWithDeps(techniqueRepo)
	getTechniqueByIDUC := application_technique.NewGetTechniqueByIDUseCase(techniqueRepo)
	getAllTechniquesUC := application_technique.NewGetAllTechniquesUseCase(techniqueRepo)
	getTechniquesListUC := application_technique.NewGetTechniquesListUseCase(techniqueRepo)
	updateTechniqueUC := application_technique.NewUpdateTechniqueUseCase(techniqueRepo)

	return presentation_technique.NewTechniqueHandler(
		createTechniqueUC,
		updateTechniqueUC,
		getTechniqueByIDUC,
		getAllTechniquesUC,
		getTechniquesListUC,
	)
}

func main() {
	router := gin.Default()

	// CORS: allow frontend at localhost:8081 (and preflight)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// User routes
	createUserUC, updateUserByIDUC, loginUserUC, getUserByIDUC, getAllUsersUC := initUserCases()
	userHandler := initUserHandler(createUserUC, updateUserByIDUC, loginUserUC, getUserByIDUC, getAllUsersUC)
	presentation_user.UserRoutes(router.Group("/api"), userHandler)

	// Training routes
	trainingHandler := initTrainingHandler()
	presentation_training.TrainingRoutes(router.Group("/api"), trainingHandler)

	// Technique routes
	techniqueHandler := initTechniqueHandler()
	presentation_technique.TechniqueRoutes(router.Group("/api"), techniqueHandler)

	router.Run()
}
