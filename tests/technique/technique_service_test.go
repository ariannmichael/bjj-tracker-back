package technique_test

import (
	"errors"
	"testing"

	application_technique "bjj-tracker/src/modules/technique/application"
	domain_technique "bjj-tracker/src/modules/technique/domain"
)

// MockTechniqueRepository is a mock implementation of TechniqueRepository
type MockTechniqueRepository struct {
	techniques      map[string]*domain_technique.Technique
	allTechniques   []domain_technique.Technique
	createError     error
	updateError     error
	findByIDError   error
	findByIDsError  error
	findAllError    error
	findByCategoryError error
	findByDifficultyError error
}

func NewMockTechniqueRepository() *MockTechniqueRepository {
	return &MockTechniqueRepository{
		techniques: make(map[string]*domain_technique.Technique),
		allTechniques: []domain_technique.Technique{},
	}
}

func (m *MockTechniqueRepository) Create(technique *domain_technique.Technique) (*domain_technique.Technique, error) {
	if m.createError != nil {
		return nil, m.createError
	}
	if technique.ID == "" {
		technique.ID = "test-id-1"
	}
	m.techniques[technique.ID] = technique
	return technique, nil
}

func (m *MockTechniqueRepository) Update(technique *domain_technique.Technique) (*domain_technique.Technique, error) {
	if m.updateError != nil {
		return nil, m.updateError
	}
	m.techniques[technique.ID] = technique
	return technique, nil
}

func (m *MockTechniqueRepository) FindByID(id string) (*domain_technique.Technique, error) {
	if m.findByIDError != nil {
		return nil, m.findByIDError
	}
	technique, exists := m.techniques[id]
	if !exists {
		return nil, errors.New("technique not found")
	}
	return technique, nil
}

func (m *MockTechniqueRepository) FindByIDs(ids []string) ([]domain_technique.Technique, error) {
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

func (m *MockTechniqueRepository) FindByCategory(category domain_technique.Category) ([]domain_technique.Technique, error) {
	if m.findByCategoryError != nil {
		return nil, m.findByCategoryError
	}
	var result []domain_technique.Technique
	for _, technique := range m.techniques {
		if technique.Category == category {
			result = append(result, *technique)
		}
	}
	return result, nil
}

func (m *MockTechniqueRepository) FindByDifficulty(difficulty domain_technique.Difficulty) ([]domain_technique.Technique, error) {
	if m.findByDifficultyError != nil {
		return nil, m.findByDifficultyError
	}
	var result []domain_technique.Technique
	for _, technique := range m.techniques {
		if technique.Difficulty == difficulty {
			result = append(result, *technique)
		}
	}
	return result, nil
}

func (m *MockTechniqueRepository) FindAll() ([]domain_technique.Technique, error) {
	if m.findAllError != nil {
		return nil, m.findAllError
	}
	if len(m.allTechniques) > 0 {
		return m.allTechniques, nil
	}
	var result []domain_technique.Technique
	for _, technique := range m.techniques {
		result = append(result, *technique)
	}
	return result, nil
}

func TestTechniqueService_CreateTechnique(t *testing.T) {
	tests := []struct {
		name      string
		request   *application_technique.CreateTechniqueRequest
		setupMock func(*MockTechniqueRepository)
		wantError bool
	}{
		{
			name: "successful creation",
			request: &application_technique.CreateTechniqueRequest{
				Name:                  "Armbar",
				NamePortuguese:        "Chave de Braço",
				Description:           "A submission technique",
				DescriptionPortuguese: "Uma técnica de finalização",
				Category:              domain_technique.Submission,
				Difficulty:            domain_technique.Beginner,
			},
			setupMock: func(m *MockTechniqueRepository) {},
			wantError: false,
		},
		{
			name: "repository error",
			request: &application_technique.CreateTechniqueRequest{
				Name:                  "Armbar",
				NamePortuguese:        "Chave de Braço",
				Description:           "A submission technique",
				DescriptionPortuguese: "Uma técnica de finalização",
				Category:              domain_technique.Submission,
				Difficulty:            domain_technique.Beginner,
			},
			setupMock: func(m *MockTechniqueRepository) {
				m.createError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTechniqueRepository()
			tt.setupMock(mockRepo)
			service := application_technique.NewTechniqueService(mockRepo)

			result, err := service.CreateTechnique(tt.request)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				if result != nil {
					t.Errorf("expected nil result but got %v", result)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result == nil {
					t.Errorf("expected result but got nil")
				}
				if result != nil && result.Name != tt.request.Name {
					t.Errorf("expected name %s but got %s", tt.request.Name, result.Name)
				}
			}
		})
	}
}

func TestTechniqueService_UpdateTechnique(t *testing.T) {
	tests := []struct {
		name      string
		request   *application_technique.UpdateTechniqueRequest
		setupMock func(*MockTechniqueRepository)
		wantError bool
	}{
		{
			name: "successful update",
			request: &application_technique.UpdateTechniqueRequest{
				ID:                    "test-id",
				Name:                  "Updated Armbar",
				NamePortuguese:        "Chave de Braço Atualizada",
				Description:           "Updated description",
				DescriptionPortuguese: "Descrição atualizada",
				Category:              domain_technique.Submission,
				Difficulty:            domain_technique.Intermediate,
			},
			setupMock: func(m *MockTechniqueRepository) {
				m.techniques["test-id"] = &domain_technique.Technique{
					ID:                    "test-id",
					Name:                  "Armbar",
					NamePortuguese:        "Chave de Braço",
					Description:           "A submission technique",
					DescriptionPortuguese: "Uma técnica de finalização",
					Category:              domain_technique.Submission,
					Difficulty:            domain_technique.Beginner,
				}
			},
			wantError: false,
		},
		{
			name: "repository error",
			request: &application_technique.UpdateTechniqueRequest{
				ID:                    "test-id",
				Name:                  "Updated Armbar",
				NamePortuguese:        "Chave de Braço Atualizada",
				Description:           "Updated description",
				DescriptionPortuguese: "Descrição atualizada",
				Category:              domain_technique.Submission,
				Difficulty:            domain_technique.Intermediate,
			},
			setupMock: func(m *MockTechniqueRepository) {
				m.updateError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTechniqueRepository()
			tt.setupMock(mockRepo)
			service := application_technique.NewTechniqueService(mockRepo)

			result, err := service.UpdateTechnique(tt.request)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != nil && result.Name != tt.request.Name {
					t.Errorf("expected name %s but got %s", tt.request.Name, result.Name)
				}
			}
		})
	}
}

func TestTechniqueService_GetTechniqueByID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		setupMock func(*MockTechniqueRepository)
		wantError bool
	}{
		{
			name: "successful retrieval",
			id:   "test-id",
			setupMock: func(m *MockTechniqueRepository) {
				m.techniques["test-id"] = &domain_technique.Technique{
					ID:     "test-id",
					Name:   "Armbar",
					Category: domain_technique.Submission,
				}
			},
			wantError: false,
		},
		{
			name: "not found",
			id:   "non-existent",
			setupMock: func(m *MockTechniqueRepository) {},
			wantError: true,
		},
		{
			name: "repository error",
			id:   "test-id",
			setupMock: func(m *MockTechniqueRepository) {
				m.findByIDError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTechniqueRepository()
			tt.setupMock(mockRepo)
			service := application_technique.NewTechniqueService(mockRepo)

			result, err := service.GetTechniqueByID(tt.id)

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

func TestTechniqueService_GetAllTechniques(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(*MockTechniqueRepository)
		wantCount int
		wantError bool
	}{
		{
			name: "successful retrieval",
			setupMock: func(m *MockTechniqueRepository) {
				m.techniques["1"] = &domain_technique.Technique{ID: "1", Name: "Armbar"}
				m.techniques["2"] = &domain_technique.Technique{ID: "2", Name: "Triangle"}
			},
			wantCount: 2,
			wantError: false,
		},
		{
			name: "empty list",
			setupMock: func(m *MockTechniqueRepository) {},
			wantCount: 0,
			wantError: false,
		},
		{
			name: "repository error",
			setupMock: func(m *MockTechniqueRepository) {
				m.findAllError = errors.New("database error")
			},
			wantCount: 0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTechniqueRepository()
			tt.setupMock(mockRepo)
			service := application_technique.NewTechniqueService(mockRepo)

			result, err := service.GetAllTechniques()

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if len(result) != tt.wantCount {
					t.Errorf("expected %d techniques but got %d", tt.wantCount, len(result))
				}
			}
		})
	}
}

func TestTechniqueService_GetTechniquesByCategory(t *testing.T) {
	tests := []struct {
		name      string
		category  domain_technique.Category
		setupMock func(*MockTechniqueRepository)
		wantCount int
		wantError bool
	}{
		{
			name:     "successful retrieval",
			category: domain_technique.Submission,
			setupMock: func(m *MockTechniqueRepository) {
				m.techniques["1"] = &domain_technique.Technique{ID: "1", Name: "Armbar", Category: domain_technique.Submission}
				m.techniques["2"] = &domain_technique.Technique{ID: "2", Name: "Sweep", Category: domain_technique.Sweep}
			},
			wantCount: 1,
			wantError: false,
		},
		{
			name:     "repository error",
			category: domain_technique.Submission,
			setupMock: func(m *MockTechniqueRepository) {
				m.findByCategoryError = errors.New("database error")
			},
			wantCount: 0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTechniqueRepository()
			tt.setupMock(mockRepo)
			service := application_technique.NewTechniqueService(mockRepo)

			result, err := service.GetTechniquesByCategory(tt.category)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if len(result) != tt.wantCount {
					t.Errorf("expected %d techniques but got %d", tt.wantCount, len(result))
				}
			}
		})
	}
}

func TestTechniqueService_GetTechniquesByDifficulty(t *testing.T) {
	tests := []struct {
		name      string
		difficulty domain_technique.Difficulty
		setupMock func(*MockTechniqueRepository)
		wantCount int
		wantError bool
	}{
		{
			name:      "successful retrieval",
			difficulty: domain_technique.Beginner,
			setupMock: func(m *MockTechniqueRepository) {
				m.techniques["1"] = &domain_technique.Technique{ID: "1", Name: "Armbar", Difficulty: domain_technique.Beginner}
				m.techniques["2"] = &domain_technique.Technique{ID: "2", Name: "Triangle", Difficulty: domain_technique.Advanced}
			},
			wantCount: 1,
			wantError: false,
		},
		{
			name:      "repository error",
			difficulty: domain_technique.Beginner,
			setupMock: func(m *MockTechniqueRepository) {
				m.findByDifficultyError = errors.New("database error")
			},
			wantCount: 0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTechniqueRepository()
			tt.setupMock(mockRepo)
			service := application_technique.NewTechniqueService(mockRepo)

			result, err := service.GetTechniquesByDifficulty(tt.difficulty)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if len(result) != tt.wantCount {
					t.Errorf("expected %d techniques but got %d", tt.wantCount, len(result))
				}
			}
		})
	}
}

func TestTechniqueService_GetTechniquesByIDs(t *testing.T) {
	tests := []struct {
		name      string
		ids       []string
		setupMock func(*MockTechniqueRepository)
		wantCount int
		wantError bool
	}{
		{
			name: "successful retrieval",
			ids:  []string{"1", "2"},
			setupMock: func(m *MockTechniqueRepository) {
				m.techniques["1"] = &domain_technique.Technique{ID: "1", Name: "Armbar"}
				m.techniques["2"] = &domain_technique.Technique{ID: "2", Name: "Triangle"}
			},
			wantCount: 2,
			wantError: false,
		},
		{
			name: "partial match",
			ids:  []string{"1", "3"},
			setupMock: func(m *MockTechniqueRepository) {
				m.techniques["1"] = &domain_technique.Technique{ID: "1", Name: "Armbar"}
				m.techniques["2"] = &domain_technique.Technique{ID: "2", Name: "Triangle"}
			},
			wantCount: 1,
			wantError: false,
		},
		{
			name: "repository error",
			ids:  []string{"1"},
			setupMock: func(m *MockTechniqueRepository) {
				m.findByIDsError = errors.New("database error")
			},
			wantCount: 0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTechniqueRepository()
			tt.setupMock(mockRepo)
			service := application_technique.NewTechniqueService(mockRepo)

			result, err := service.GetTechniquesByIDs(tt.ids)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if len(result) != tt.wantCount {
					t.Errorf("expected %d techniques but got %d", tt.wantCount, len(result))
				}
			}
		})
	}
}

