package training_test

import (
	"errors"
	"testing"

	application_technique "bjj-tracker/src/modules/technique/application"
	application_training "bjj-tracker/src/modules/training/application"
	domain_technique "bjj-tracker/src/modules/technique/domain"
	domain_training "bjj-tracker/src/modules/training/domain"
)

// MockTechniqueRepositoryForTraining is a mock technique repository for training tests
type MockTechniqueRepositoryForTraining struct {
	techniques    map[string]*domain_technique.Technique
	findByIDsError error
}

func NewMockTechniqueRepositoryForTraining() *MockTechniqueRepositoryForTraining {
	return &MockTechniqueRepositoryForTraining{
		techniques: make(map[string]*domain_technique.Technique),
	}
}

func (m *MockTechniqueRepositoryForTraining) Create(technique *domain_technique.Technique) (*domain_technique.Technique, error) {
	return nil, nil
}

func (m *MockTechniqueRepositoryForTraining) Update(technique *domain_technique.Technique) (*domain_technique.Technique, error) {
	return nil, nil
}

func (m *MockTechniqueRepositoryForTraining) FindByID(id string) (*domain_technique.Technique, error) {
	return nil, nil
}

func (m *MockTechniqueRepositoryForTraining) FindByIDs(ids []string) ([]domain_technique.Technique, error) {
	if m.findByIDsError != nil {
		return nil, m.findByIDsError
	}
	var result []domain_technique.Technique
	for _, id := range ids {
		if technique, exists := m.techniques[id]; exists {
			result = append(result, *technique)
		}
	}
	return result, nil
}

func (m *MockTechniqueRepositoryForTraining) FindByCategory(category domain_technique.Category) ([]domain_technique.Technique, error) {
	return nil, nil
}

func (m *MockTechniqueRepositoryForTraining) FindByDifficulty(difficulty domain_technique.Difficulty) ([]domain_technique.Technique, error) {
	return nil, nil
}

func (m *MockTechniqueRepositoryForTraining) FindAll() ([]domain_technique.Technique, error) {
	return nil, nil
}

func TestCreateTrainingUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		request   application_training.CreateTrainingRequest
		setupMock func(*MockTrainingRepository, *MockTechniqueRepositoryForTraining)
		wantError bool
	}{
		{
			name: "successful creation",
			request: application_training.CreateTrainingRequest{
				UserID:       "test-user-id",
				TechniqueIDs: []string{"tech-1", "tech-2"},
				Duration:     60,
				Notes:        "Great training session",
			},
			setupMock: func(tr *MockTrainingRepository, trr *MockTechniqueRepositoryForTraining) {
				trr.techniques["tech-1"] = &domain_technique.Technique{ID: "tech-1", Name: "Armbar"}
				trr.techniques["tech-2"] = &domain_technique.Technique{ID: "tech-2", Name: "Triangle"}
			},
			wantError: false,
		},
		{
			name: "technique service error",
			request: application_training.CreateTrainingRequest{
				UserID:       "test-user-id",
				TechniqueIDs: []string{"tech-1"},
				Duration:     60,
				Notes:        "Great training session",
			},
			setupMock: func(tr *MockTrainingRepository, trr *MockTechniqueRepositoryForTraining) {
				trr.findByIDsError = errors.New("technique not found")
			},
			wantError: true,
		},
		{
			name: "repository error",
			request: application_training.CreateTrainingRequest{
				UserID:       "test-user-id",
				TechniqueIDs: []string{"tech-1"},
				Duration:     60,
				Notes:        "Great training session",
			},
			setupMock: func(tr *MockTrainingRepository, trr *MockTechniqueRepositoryForTraining) {
				trr.techniques["tech-1"] = &domain_technique.Technique{ID: "tech-1", Name: "Armbar"}
				tr.createError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTrainingRepository()
			mockTechniqueRepo := NewMockTechniqueRepositoryForTraining()
			tt.setupMock(mockRepo, mockTechniqueRepo)
			
			techniqueService := application_technique.NewTechniqueService(mockTechniqueRepo)
			useCase := &application_training.CreateTrainingUseCase{
				Repo:             mockRepo,
				TechniqueService: techniqueService,
			}

			result, err := useCase.Execute(tt.request)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result == nil {
					t.Errorf("expected result but got nil")
				}
				if result != nil && result.UserID != tt.request.UserID {
					t.Errorf("expected UserID %s but got %s", tt.request.UserID, result.UserID)
				}
			}
		})
	}
}

func TestGetTrainingByIDUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		setupMock func(*MockTrainingRepository)
		wantError bool
	}{
		{
			name: "successful retrieval",
			id:   "test-training-id",
			setupMock: func(m *MockTrainingRepository) {
				m.trainingSessions["test-training-id"] = &domain_training.TrainingSession{
					ID:       "test-training-id",
					UserID:   "test-user-id",
					Duration: 60,
					Notes:    "Great session",
				}
			},
			wantError: false,
		},
		{
			name:      "not found",
			id:        "non-existent",
			setupMock: func(m *MockTrainingRepository) {},
			wantError: true,
		},
		{
			name: "repository error",
			id:   "test-training-id",
			setupMock: func(m *MockTrainingRepository) {
				m.getByIDError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTrainingRepository()
			tt.setupMock(mockRepo)
			useCase := application_training.NewGetTrainingByIDUseCase(mockRepo)

			result, err := useCase.Execute(tt.id)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result == nil {
					t.Errorf("expected result but got nil")
				}
				if result != nil && result.ID != tt.id {
					t.Errorf("expected ID %s but got %s", tt.id, result.ID)
				}
			}
		})
	}
}

func TestGetAllTrainingsUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(*MockTrainingRepository)
		wantCount int
		wantError bool
	}{
		{
			name: "successful retrieval",
			setupMock: func(m *MockTrainingRepository) {
				m.trainingSessions["1"] = &domain_training.TrainingSession{
					ID:       "1",
					UserID:   "user-1",
					Duration: 60,
				}
				m.trainingSessions["2"] = &domain_training.TrainingSession{
					ID:       "2",
					UserID:   "user-2",
					Duration: 90,
				}
			},
			wantCount: 2,
			wantError: false,
		},
		{
			name:      "empty list",
			setupMock: func(m *MockTrainingRepository) {},
			wantCount: 0,
			wantError: false,
		},
		{
			name: "repository error",
			setupMock: func(m *MockTrainingRepository) {
				m.getAllError = errors.New("database error")
			},
			wantCount: 0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTrainingRepository()
			tt.setupMock(mockRepo)
			useCase := application_training.NewGetAllTrainingsUseCase(mockRepo)

			result, err := useCase.Execute()

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if len(result) != tt.wantCount {
					t.Errorf("expected %d training sessions but got %d", tt.wantCount, len(result))
				}
			}
		})
	}
}

func TestUpdateTrainingUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		request   application_training.UpdateTrainingRequest
		setupMock func(*MockTrainingRepository, *MockTechniqueRepositoryForTraining)
		wantError bool
	}{
		{
			name: "successful update",
			id:   "test-training-id",
			request: application_training.UpdateTrainingRequest{
				TechniqueIDs: []string{"tech-1", "tech-2"},
				Duration:     90,
				Notes:        "Updated notes",
			},
			setupMock: func(tr *MockTrainingRepository, trr *MockTechniqueRepositoryForTraining) {
				tr.trainingSessions["test-training-id"] = &domain_training.TrainingSession{
					ID:       "test-training-id",
					UserID:   "test-user-id",
					Duration: 60,
					Notes:    "Old notes",
				}
				trr.techniques["tech-1"] = &domain_technique.Technique{ID: "tech-1", Name: "Armbar"}
				trr.techniques["tech-2"] = &domain_technique.Technique{ID: "tech-2", Name: "Triangle"}
			},
			wantError: false,
		},
		{
			name: "training not found",
			id:   "non-existent",
			request: application_training.UpdateTrainingRequest{
				TechniqueIDs: []string{"tech-1"},
				Duration:     90,
				Notes:        "Updated notes",
			},
			setupMock: func(tr *MockTrainingRepository, trr *MockTechniqueRepositoryForTraining) {},
			wantError: true,
		},
		{
			name: "technique service error",
			id:   "test-training-id",
			request: application_training.UpdateTrainingRequest{
				TechniqueIDs: []string{"tech-1"},
				Duration:     90,
				Notes:        "Updated notes",
			},
			setupMock: func(tr *MockTrainingRepository, trr *MockTechniqueRepositoryForTraining) {
				tr.trainingSessions["test-training-id"] = &domain_training.TrainingSession{
					ID:       "test-training-id",
					UserID:   "test-user-id",
					Duration: 60,
				}
				trr.findByIDsError = errors.New("technique not found")
			},
			wantError: true,
		},
		{
			name: "repository update error",
			id:   "test-training-id",
			request: application_training.UpdateTrainingRequest{
				TechniqueIDs: []string{"tech-1"},
				Duration:     90,
				Notes:        "Updated notes",
			},
			setupMock: func(tr *MockTrainingRepository, trr *MockTechniqueRepositoryForTraining) {
				tr.trainingSessions["test-training-id"] = &domain_training.TrainingSession{
					ID:       "test-training-id",
					UserID:   "test-user-id",
					Duration: 60,
				}
				trr.techniques["tech-1"] = &domain_technique.Technique{ID: "tech-1", Name: "Armbar"}
				tr.updateError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTrainingRepository()
			mockTechniqueRepo := NewMockTechniqueRepositoryForTraining()
			tt.setupMock(mockRepo, mockTechniqueRepo)
			
			techniqueService := application_technique.NewTechniqueService(mockTechniqueRepo)
			useCase := &application_training.UpdateTrainingUseCase{
				Repo:             mockRepo,
				TechniqueService: techniqueService,
			}

			result, err := useCase.Execute(tt.id, tt.request)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result == nil {
					t.Errorf("expected result but got nil")
				}
				if result != nil && result.Duration != tt.request.Duration {
					t.Errorf("expected Duration %d but got %d", tt.request.Duration, result.Duration)
				}
			}
		})
	}
}

func TestDeleteTrainingUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		setupMock func(*MockTrainingRepository)
		wantError bool
	}{
		{
			name: "successful deletion",
			id:   "test-training-id",
			setupMock: func(m *MockTrainingRepository) {
				m.trainingSessions["test-training-id"] = &domain_training.TrainingSession{
					ID:       "test-training-id",
					UserID:   "test-user-id",
					Duration: 60,
					Notes:    "Great session",
				}
			},
			wantError: false,
		},
		{
			name:      "training not found",
			id:        "non-existent",
			setupMock: func(m *MockTrainingRepository) {},
			wantError: false, // Delete typically doesn't error if not found
		},
		{
			name: "repository error",
			id:   "test-training-id",
			setupMock: func(m *MockTrainingRepository) {
				m.trainingSessions["test-training-id"] = &domain_training.TrainingSession{
					ID:       "test-training-id",
					UserID:   "test-user-id",
					Duration: 60,
				}
				m.deleteError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTrainingRepository()
			tt.setupMock(mockRepo)
			useCase := application_training.NewDeleteTrainingUseCase(mockRepo)

			err := useCase.Execute(tt.id)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

