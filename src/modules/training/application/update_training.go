package application_training

import (
	application_technique "jiu-tracker/src/modules/technique/application"
	domain_training "jiu-tracker/src/modules/training/domain"
	"fmt"
)

type UpdateTrainingUseCase struct {
	Repo             domain_training.TrainingRepository
	TechniqueService *application_technique.TechniqueService
}

func NewUpdateTrainingUseCase(repo domain_training.TrainingRepository, techniqueService *application_technique.TechniqueService) *UpdateTrainingUseCase {
	return &UpdateTrainingUseCase{Repo: repo, TechniqueService: techniqueService}
}

func (uc *UpdateTrainingUseCase) Execute(id string, req UpdateTrainingRequest) (*domain_training.TrainingSession, error) {
	training, err := uc.Repo.GetTrainingSessionByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find training: %w", err)
	}
	submitUsing, err := uc.TechniqueService.GetTechniquesByIDs(req.SubmitUsingOptionsIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to find submit-using techniques: %w", err)
	}
	submittedBy, err := uc.TechniqueService.GetTechniquesByIDs(req.SubmittedByOptionsIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to find submitted-by techniques: %w", err)
	}
	training.Date = req.Date
	training.IsOpenMat = req.IsOpenMat
	training.SubmitUsingOptions = submitUsing
	training.SubmittedByOptions = submittedBy
	training.Duration = req.Duration
	training.Notes = req.Notes
	newTraining, err := uc.Repo.UpdateTrainingSession(training)
	if err != nil {
		return nil, fmt.Errorf("failed to update training: %w", err)
	}
	if newTraining == nil {
		return nil, fmt.Errorf("failed to update training: training is nil")
	}
	return newTraining, nil
}
