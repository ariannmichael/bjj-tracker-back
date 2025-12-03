package application_training

import (
	domain_training "bjj-tracker/src/modules/training/domain"
	"fmt"
)

type DeleteTrainingUseCase struct {
	Repo domain_training.TrainingRepository
}

func NewDeleteTrainingUseCase(repo domain_training.TrainingRepository) *DeleteTrainingUseCase {
	return &DeleteTrainingUseCase{Repo: repo}
}

func (uc *DeleteTrainingUseCase) Execute(id string) error {
	err := uc.Repo.DeleteTrainingSession(id)
	if err != nil {
		return fmt.Errorf("failed to delete training: %w", err)
	}
	return nil
}

