package application_training

import (
	"bjj-tracker/config"
	application_technique "bjj-tracker/src/modules/technique/application"
	infrastructure_technique "bjj-tracker/src/modules/technique/infrastructure"
	domain_training "bjj-tracker/src/modules/training/domain"
	infrastructure_training "bjj-tracker/src/modules/training/infrastructure"
)

type CreateTrainingUseCase struct {
	Repo             domain_training.TrainingRepository
	TechniqueService *application_technique.TechniqueService
}

func NewCreateTrainingUseCase() *CreateTrainingUseCase {
	db := config.ConnectToDB()
	trainingRepo := &infrastructure_training.TrainingRepositoryImpl{DB: db}
	techniqueRepo := &infrastructure_technique.TechniqueRepositoryImpl{DB: db}
	techniqueService := application_technique.NewTechniqueService(techniqueRepo)

	return &CreateTrainingUseCase{
		Repo:             trainingRepo,
		TechniqueService: techniqueService,
	}
}

func (uc *CreateTrainingUseCase) Execute(req CreateTrainingRequest) (*domain_training.TrainingSession, error) {
	techniques, err := uc.TechniqueService.GetTechniquesByIDs(req.TechniqueIDs)
	if err != nil {
		return nil, err
	}
	trainingSession := domain_training.TrainingSession{
		UserID:     req.UserID,
		Techniques: techniques,
		Duration:   req.Duration,
		Notes:      req.Notes,
	}
	return uc.Repo.CreateTrainingSession(&trainingSession)
}
