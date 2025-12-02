package application_training

import (
	application_technique "bjj-tracker/src/modules/technique/application"
	domain_training "bjj-tracker/src/modules/training/domain"
	"fmt"
)

type UpdateTrainingUseCase struct {
	Repo             domain_training.TrainingRepository
	TechniqueService *application_technique.TechniqueService
}

func NewUpdateTrainingUseCase(repo domain_training.TrainingRepository) *UpdateTrainingUseCase {
	return &UpdateTrainingUseCase{Repo: repo}
}

func (uc *UpdateTrainingUseCase) Execute(id string, req UpdateTrainingRequest) (*domain_training.TrainingSession, error) {
	training, err := uc.Repo.GetTrainingSessionByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find training: %w", err)
	}
	techniques, err := uc.TechniqueService.GetTechniquesByIDs(req.TechniqueIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to find techniques: %w", err)
	}
	training.Techniques = techniques
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
