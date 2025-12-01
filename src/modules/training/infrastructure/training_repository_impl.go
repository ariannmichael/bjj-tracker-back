package infrastructure_training

import (
	domain_training "bjj-tracker/src/modules/training/domain"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type TrainingRepositoryImpl struct {
	DB *gorm.DB
}

var _ domain_training.TrainingRepository = &TrainingRepositoryImpl{}

func (r *TrainingRepositoryImpl) CreateTrainingSession(trainingSession *domain_training.TrainingSession) (*domain_training.TrainingSession, error) {
	trainingSession.ID = uuid.New().String()

	if err := r.DB.Create(trainingSession).Error; err != nil {
		return nil, err
	}
	return trainingSession, nil
}

func (r *TrainingRepositoryImpl) GetTrainingSessionByID(id string) (*domain_training.TrainingSession, error) {
	var trainingSession domain_training.TrainingSession
	if err := r.DB.First(&trainingSession, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &trainingSession, nil
}

func (r *TrainingRepositoryImpl) GetAllTrainingSessions() ([]domain_training.TrainingSession, error) {
	var trainingSessions []domain_training.TrainingSession
	if err := r.DB.Find(&trainingSessions).Error; err != nil {
		return nil, err
	}
	return trainingSessions, nil
}

func (r *TrainingRepositoryImpl) UpdateTrainingSession(trainingSession *domain_training.TrainingSession) (*domain_training.TrainingSession, error) {
	if err := r.DB.Save(trainingSession).Error; err != nil {
		return nil, err
	}
	return trainingSession, nil
}

func (r *TrainingRepositoryImpl) DeleteTrainingSession(id string) error {
	if err := r.DB.Delete(&domain_training.TrainingSession{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
