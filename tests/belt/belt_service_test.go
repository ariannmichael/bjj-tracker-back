package belt_test

import (
	"errors"
	"testing"

	application_belt "bjj-tracker/src/modules/belt/application"
	domain_belt "bjj-tracker/src/modules/belt/domain"
)

// MockBeltProgressRepository is a mock implementation of BeltProgressRepository
type MockBeltProgressRepository struct {
	beltProgresses map[int]*domain_belt.BeltProgress
	createError    error
	nextID         int
}

func NewMockBeltProgressRepository() *MockBeltProgressRepository {
	return &MockBeltProgressRepository{
		beltProgresses: make(map[int]*domain_belt.BeltProgress),
		nextID:         1,
	}
}

func (m *MockBeltProgressRepository) CreateBeltProgress(beltProgress *domain_belt.BeltProgress) (*domain_belt.BeltProgress, error) {
	if m.createError != nil {
		return nil, m.createError
	}
	if beltProgress.ID == 0 {
		beltProgress.ID = m.nextID
		m.nextID++
	}
	m.beltProgresses[beltProgress.ID] = beltProgress
	return beltProgress, nil
}

func TestBeltService_GetBeltByColor(t *testing.T) {
	tests := []struct {
		name      string
		color     string
		wantBelt  domain_belt.Belt
		wantError bool
	}{
		{
			name:      "white belt",
			color:     "white",
			wantBelt:  domain_belt.White,
			wantError: false,
		},
		{
			name:      "blue belt",
			color:     "blue",
			wantBelt:  domain_belt.Blue,
			wantError: false,
		},
		{
			name:      "purple belt",
			color:     "purple",
			wantBelt:  domain_belt.Purple,
			wantError: false,
		},
		{
			name:      "brown belt",
			color:     "brown",
			wantBelt:  domain_belt.Brown,
			wantError: false,
		},
		{
			name:      "black belt",
			color:     "black",
			wantBelt:  domain_belt.Black,
			wantError: false,
		},
		{
			name:      "case insensitive - WHITE",
			color:     "WHITE",
			wantBelt:  domain_belt.White,
			wantError: false,
		},
		{
			name:      "case insensitive - Blue",
			color:     "Blue",
			wantBelt:  domain_belt.Blue,
			wantError: false,
		},
		{
			name:      "invalid color",
			color:     "red",
			wantBelt:  0,
			wantError: true,
		},
		{
			name:      "empty color",
			color:     "",
			wantBelt:  0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockBeltProgressRepository()
			service := application_belt.NewBeltService(mockRepo)

			result, err := service.GetBeltByColor(tt.color)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.wantBelt {
					t.Errorf("expected belt %v but got %v", tt.wantBelt, result)
				}
			}
		})
	}
}

func TestBeltService_CreateBeltProgress(t *testing.T) {
	tests := []struct {
		name      string
		dto       application_belt.CreateBeltProgressDTO
		setupMock func(*MockBeltProgressRepository)
		wantError bool
	}{
		{
			name: "successful creation - white belt",
			dto: application_belt.CreateBeltProgressDTO{
				UserID:  "test-user-id",
				Color:   "white",
				Stripes: 0,
			},
			setupMock: func(m *MockBeltProgressRepository) {},
			wantError: false,
		},
		{
			name: "successful creation - blue belt with stripes",
			dto: application_belt.CreateBeltProgressDTO{
				UserID:  "test-user-id",
				Color:   "blue",
				Stripes: 2,
			},
			setupMock: func(m *MockBeltProgressRepository) {},
			wantError: false,
		},
		{
			name: "invalid belt color",
			dto: application_belt.CreateBeltProgressDTO{
				UserID:  "test-user-id",
				Color:   "red",
				Stripes: 0,
			},
			setupMock: func(m *MockBeltProgressRepository) {},
			wantError: true,
		},
		{
			name: "repository error",
			dto: application_belt.CreateBeltProgressDTO{
				UserID:  "test-user-id",
				Color:   "white",
				Stripes: 0,
			},
			setupMock: func(m *MockBeltProgressRepository) {
				m.createError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockBeltProgressRepository()
			tt.setupMock(mockRepo)
			service := application_belt.NewBeltService(mockRepo)

			result, err := service.CreateBeltProgress(tt.dto)

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
				if result != nil {
					if result.UserID != tt.dto.UserID {
						t.Errorf("expected UserID %s but got %s", tt.dto.UserID, result.UserID)
					}
					if result.StripeCount != tt.dto.Stripes {
						t.Errorf("expected StripeCount %d but got %d", tt.dto.Stripes, result.StripeCount)
					}
				}
			}
		})
	}
}

func TestCreateBeltProgressUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		dto       application_belt.CreateBeltProgressDTO
		setupMock func(*MockBeltProgressRepository)
		wantError bool
	}{
		{
			name: "successful creation",
			dto: application_belt.CreateBeltProgressDTO{
				UserID:  "test-user-id",
				Color:   "blue",
				Stripes: 2,
			},
			setupMock: func(m *MockBeltProgressRepository) {},
			wantError: false,
		},
		{
			name: "invalid color",
			dto: application_belt.CreateBeltProgressDTO{
				UserID:  "test-user-id",
				Color:   "invalid",
				Stripes: 0,
			},
			setupMock: func(m *MockBeltProgressRepository) {},
			wantError: true,
		},
		{
			name: "repository error",
			dto: application_belt.CreateBeltProgressDTO{
				UserID:  "test-user-id",
				Color:   "white",
				Stripes: 0,
			},
			setupMock: func(m *MockBeltProgressRepository) {
				m.createError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockBeltProgressRepository()
			tt.setupMock(mockRepo)
			service := application_belt.NewBeltService(mockRepo)
			useCase := &application_belt.CreateBeltProgressUseCase{
				Repo:        mockRepo,
				BeltService: *service,
			}

			result, err := useCase.Execute(tt.dto)

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
				if result != nil && result.UserID != tt.dto.UserID {
					t.Errorf("expected UserID %s but got %s", tt.dto.UserID, result.UserID)
				}
			}
		})
	}
}

