package application_training

import domain_training "bjj-tracker/src/modules/training/domain"

type GetTrainingByIDUseCase struct {
	Repo domain_training.TrainingRepository
}

func NewGetTrainingByIDUseCase(repo domain_training.TrainingRepository) *GetTrainingByIDUseCase {
	return &GetTrainingByIDUseCase{Repo: repo}
}

func (uc *GetTrainingByIDUseCase) Execute(id string) (*domain_training.TrainingSession, error) {
	return uc.Repo.GetTrainingSessionByID(id)
}
