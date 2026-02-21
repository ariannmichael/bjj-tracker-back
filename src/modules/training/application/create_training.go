package application_training

import (
	"jiu-tracker/config"
	application_technique "jiu-tracker/src/modules/technique/application"
	infrastructure_technique "jiu-tracker/src/modules/technique/infrastructure"
	domain_training "jiu-tracker/src/modules/training/domain"
	infrastructure_training "jiu-tracker/src/modules/training/infrastructure"
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
	return NewCreateTrainingUseCaseWithDeps(trainingRepo, techniqueService)
}

func NewCreateTrainingUseCaseWithDeps(repo domain_training.TrainingRepository, techniqueService *application_technique.TechniqueService) *CreateTrainingUseCase {
	return &CreateTrainingUseCase{Repo: repo, TechniqueService: techniqueService}
}

func (uc *CreateTrainingUseCase) Execute(req CreateTrainingRequest) (*domain_training.TrainingSession, error) {
	submitUsing, err := uc.TechniqueService.GetTechniquesByIDs(req.SubmitUsingOptionsIDs)
	if err != nil {
		return nil, err
	}
	submittedBy, err := uc.TechniqueService.GetTechniquesByIDs(req.SubmittedByOptionsIDs)
	if err != nil {
		return nil, err
	}
	trainingSession := domain_training.TrainingSession{
		UserID:                 req.UserID,
		Date:                   req.Date,
		IsOpenMat:              req.IsOpenMat,
		SubmitUsingOptions:     submitUsing,
		SubmittedByOptions:     submittedBy,
		Duration:               req.Duration,
		Notes:                  req.Notes,
	}
	return uc.Repo.CreateTrainingSession(&trainingSession)
}
