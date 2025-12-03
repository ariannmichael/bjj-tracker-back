package technique_test

import (
	"errors"
	"testing"

	application_technique "bjj-tracker/src/modules/technique/application"
	domain_technique "bjj-tracker/src/modules/technique/domain"
)

func TestCreateTechniqueUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		request   application_technique.CreateTechniqueRequest
		setupMock func(*MockTechniqueRepository)
		wantError bool
	}{
		{
			name: "successful creation",
			request: application_technique.CreateTechniqueRequest{
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
			request: application_technique.CreateTechniqueRequest{
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
			useCase := &application_technique.CreateTechniqueUseCase{
				Repo:             mockRepo,
				TechniqueService: service,
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
				if result != nil && result.Name != tt.request.Name {
					t.Errorf("expected name %s but got %s", tt.request.Name, result.Name)
				}
			}
		})
	}
}

func TestGetTechniqueByIDUseCase_Execute(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTechniqueRepository()
			tt.setupMock(mockRepo)
			useCase := application_technique.NewGetTechniqueByIDUseCase(mockRepo)

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

func TestGetAllTechniquesUseCase_Execute(t *testing.T) {
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
			name:      "empty list",
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
			useCase := application_technique.NewGetAllTechniquesUseCase(mockRepo)

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
					t.Errorf("expected %d techniques but got %d", tt.wantCount, len(result))
				}
			}
		})
	}
}

func TestUpdateTechniqueUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		request   application_technique.UpdateTechniqueRequest
		setupMock func(*MockTechniqueRepository)
		wantError bool
	}{
		{
			name: "successful update",
			id:   "test-id",
			request: application_technique.UpdateTechniqueRequest{
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
			name: "technique not found",
			id:   "non-existent",
			request: application_technique.UpdateTechniqueRequest{
				ID:                    "non-existent",
				Name:                  "Updated Armbar",
				NamePortuguese:        "Chave de Braço Atualizada",
				Description:           "Updated description",
				DescriptionPortuguese: "Descrição atualizada",
				Category:              domain_technique.Submission,
				Difficulty:            domain_technique.Intermediate,
			},
			setupMock: func(m *MockTechniqueRepository) {},
			wantError: true,
		},
		{
			name: "update error",
			id:   "test-id",
			request: application_technique.UpdateTechniqueRequest{
				ID:                    "test-id",
				Name:                  "Updated Armbar",
				NamePortuguese:        "Chave de Braço Atualizada",
				Description:           "Updated description",
				DescriptionPortuguese: "Descrição atualizada",
				Category:              domain_technique.Submission,
				Difficulty:            domain_technique.Intermediate,
			},
			setupMock: func(m *MockTechniqueRepository) {
				m.techniques["test-id"] = &domain_technique.Technique{ID: "test-id", Name: "Armbar"}
				m.updateError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockTechniqueRepository()
			tt.setupMock(mockRepo)
			useCase := application_technique.NewUpdateTechniqueUseCase(mockRepo)

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
				if result != nil && result.Name != tt.request.Name {
					t.Errorf("expected name %s but got %s", tt.request.Name, result.Name)
				}
			}
		})
	}
}

