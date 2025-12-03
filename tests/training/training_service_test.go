package training_test

import (
	"errors"
	"testing"

	application_technique "bjj-tracker/src/modules/technique/application"
	application_training "bjj-tracker/src/modules/training/application"
	domain_technique "bjj-tracker/src/modules/technique/domain"
	domain_training "bjj-tracker/src/modules/training/domain"
)

// MockTrainingRepository is a mock implementation of TrainingRepository
type MockTrainingRepository struct {
	trainingSessions      map[string]*domain_training.TrainingSession
	createError           error
	updateError           error
	getByIDError          error
	getAllError           error
	deleteError           error
}

func NewMockTrainingRepository() *MockTrainingRepository {
	return &MockTrainingRepository{
		trainingSessions: make(map[string]*domain_training.TrainingSession),
	}
}

func (m *MockTrainingRepository) CreateTrainingSession(trainingSession *domain_training.TrainingSession) (*domain_training.TrainingSession, error) {
	if m.createError != nil {
		return nil, m.createError
	}
	if trainingSession.ID == "" {
		trainingSession.ID = "test-training-id"
	}
	m.trainingSessions[trainingSession.ID] = trainingSession
	return trainingSession, nil
}

func (m *MockTrainingRepository) GetTrainingSessionByID(id string) (*domain_training.TrainingSession, error) {
	if m.getByIDError != nil {
		return nil, m.getByIDError
	}
	training, exists := m.trainingSessions[id]
	if !exists {
		return nil, errors.New("training session not found")
	}
	return training, nil
}

func (m *MockTrainingRepository) GetAllTrainingSessions() ([]domain_training.TrainingSession, error) {
	if m.getAllError != nil {
		return nil, m.getAllError
	}
	var result []domain_training.TrainingSession
	for _, training := range m.trainingSessions {
		result = append(result, *training)
	}
	return result, nil
}

func (m *MockTrainingRepository) UpdateTrainingSession(trainingSession *domain_training.TrainingSession) (*domain_training.TrainingSession, error) {
	if m.updateError != nil {
		return nil, m.updateError
	}
	m.trainingSessions[trainingSession.ID] = trainingSession
	return trainingSession, nil
}

func (m *MockTrainingRepository) DeleteTrainingSession(id string) error {
	if m.deleteError != nil {
		return m.deleteError
	}
	delete(m.trainingSessions, id)
	return nil
}

// MockTechniqueService is a mock implementation of TechniqueService
type MockTechniqueService struct {
	techniques map[string]*domain_technique.Technique
	getByIDsError error
}

func NewMockTechniqueService() *MockTechniqueService {
	return &MockTechniqueService{
		techniques: make(map[string]*domain_technique.Technique),
	}
}

func (m *MockTechniqueService) GetTechniquesByIDs(ids []string) ([]domain_technique.Technique, error) {
	if m.getByIDsError != nil {
		return nil, m.getByIDsError
	}
	var result []domain_technique.Technique
	for _, id := range ids {
		if technique, exists := m.techniques[id]; exists {
			result = append(result, *technique)
		}
	}
	return result, nil
}

func (m *MockTechniqueService) CreateTechnique(techniqueDTO *application_technique.CreateTechniqueRequest) (*domain_technique.Technique, error) {
	return nil, nil
}

func (m *MockTechniqueService) UpdateTechnique(techniqueDTO *application_technique.UpdateTechniqueRequest) (*domain_technique.Technique, error) {
	return nil, nil
}

func (m *MockTechniqueService) GetTechniqueByID(id string) (*domain_technique.Technique, error) {
	return nil, nil
}

func (m *MockTechniqueService) GetTechniquesByCategory(category domain_technique.Category) ([]*domain_technique.Technique, error) {
	return nil, nil
}

func (m *MockTechniqueService) GetTechniquesByDifficulty(difficulty domain_technique.Difficulty) ([]*domain_technique.Technique, error) {
	return nil, nil
}

func (m *MockTechniqueService) GetAllTechniques() ([]*domain_technique.Technique, error) {
	return nil, nil
}

func TestTrainingService_NewTrainingService(t *testing.T) {
	mockRepo := NewMockTrainingRepository()
	service := application_training.NewTrainingService(mockRepo)

	if service == nil {
		t.Errorf("expected service to be created but got nil")
	}
	if service.Repo != mockRepo {
		t.Errorf("expected repo to be set correctly")
	}
}

