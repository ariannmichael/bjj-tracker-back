package application_training

import domain_training "bjj-tracker/src/modules/training/domain"

type GetAllTrainingsUseCase struct {
	Repo domain_training.TrainingRepository
}

func NewGetAllTrainingsUseCase(repo domain_training.TrainingRepository) *GetAllTrainingsUseCase {
	return &GetAllTrainingsUseCase{Repo: repo}
}

func (uc *GetAllTrainingsUseCase) Execute() ([]domain_training.TrainingSession, error) {
	return uc.Repo.GetAllTrainingSessions()
}
