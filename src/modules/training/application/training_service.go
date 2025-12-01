package application_training

import (
	application_technique "bjj-tracker/src/modules/technique/application"
	domain_training "bjj-tracker/src/modules/training/domain"
)

type TrainingService struct {
	Repo             domain_training.TrainingRepository
	TechniqueService application_technique.TechniqueService
}

func NewTrainingService(repo domain_training.TrainingRepository) *TrainingService {
	return &TrainingService{
		Repo: repo,
	}
}
