package user_test

import (
	"errors"
	"testing"

	application_belt "bjj-tracker/src/modules/belt/application"
	application_user "bjj-tracker/src/modules/user/application"
	domain_belt "bjj-tracker/src/modules/belt/domain"
	domain_user "bjj-tracker/src/modules/user/domain"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	users         map[string]*domain_user.User
	createError   error
	updateError   error
	findByIDError error
	findByEmailError error
	findAllError  error
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*domain_user.User),
	}
}

func (m *MockUserRepository) Create(user *domain_user.User) (*domain_user.User, error) {
	if m.createError != nil {
		return nil, m.createError
	}
	if user.ID == "" {
		user.ID = "test-user-id"
	}
	m.users[user.ID] = user
	return user, nil
}

func (m *MockUserRepository) Update(user *domain_user.User) (*domain_user.User, error) {
	if m.updateError != nil {
		return nil, m.updateError
	}
	m.users[user.ID] = user
	return user, nil
}

func (m *MockUserRepository) FindByID(id string) (*domain_user.User, error) {
	if m.findByIDError != nil {
		return nil, m.findByIDError
	}
	user, exists := m.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (m *MockUserRepository) FindByEmail(email string) (*domain_user.User, error) {
	if m.findByEmailError != nil {
		return nil, m.findByEmailError
	}
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *MockUserRepository) FindAll() ([]domain_user.User, error) {
	if m.findAllError != nil {
		return nil, m.findAllError
	}
	var result []domain_user.User
	for _, user := range m.users {
		result = append(result, *user)
	}
	return result, nil
}

// MockBeltService is a mock implementation of BeltService
type MockBeltService struct {
	beltProgress *domain_belt.BeltProgress
	createError  error
}

func NewMockBeltService() *MockBeltService {
	return &MockBeltService{}
}

func (m *MockBeltService) CreateBeltProgress(cbDTO application_belt.CreateBeltProgressDTO) (*domain_belt.BeltProgress, error) {
	if m.createError != nil {
		return nil, m.createError
	}
	if m.beltProgress != nil {
		return m.beltProgress, nil
	}
	return &domain_belt.BeltProgress{
		UserID:      cbDTO.UserID,
		CurrentBelt: domain_belt.White,
		StripeCount: cbDTO.Stripes,
	}, nil
}

func (m *MockBeltService) GetBeltByColor(color string) (domain_belt.Belt, error) {
	switch color {
	case "white":
		return domain_belt.White, nil
	case "blue":
		return domain_belt.Blue, nil
	case "purple":
		return domain_belt.Purple, nil
	case "brown":
		return domain_belt.Brown, nil
	case "black":
		return domain_belt.Black, nil
	default:
		return 0, errors.New("invalid belt color")
	}
}

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

func TestUserService_CreateUser(t *testing.T) {
	tests := []struct {
		name      string
		request   *application_user.CreateUserRequest
		setupMock func(*MockUserRepository)
		wantError bool
	}{
		{
			name: "successful creation",
			request: &application_user.CreateUserRequest{
				Name:     "John Doe",
				Username: "johndoe",
				Email:    "john@example.com",
				Password: "password123",
			},
			setupMock: func(m *MockUserRepository) {},
			wantError: false,
		},
		{
			name: "repository error",
			request: &application_user.CreateUserRequest{
				Name:     "John Doe",
				Username: "johndoe",
				Email:    "john@example.com",
				Password: "password123",
			},
			setupMock: func(m *MockUserRepository) {
				m.createError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository()
			tt.setupMock(mockRepo)
			mockBeltRepo := NewMockBeltProgressRepository()
			beltService := application_belt.NewBeltService(mockBeltRepo)
			service := application_user.NewUserService(mockRepo, *beltService)

			result, err := service.CreateUser(tt.request)

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

func TestUserService_UpdateUser(t *testing.T) {
	tests := []struct {
		name      string
		user      *domain_user.User
		setupMock func(*MockUserRepository)
		wantError bool
	}{
		{
			name: "successful update",
			user: &domain_user.User{
				ID:       "test-user-id",
				Name:     "Updated Name",
				Username: "updateduser",
				Email:    "updated@example.com",
				Password: "newpassword",
			},
			setupMock: func(m *MockUserRepository) {
				m.users["test-user-id"] = &domain_user.User{
					ID:       "test-user-id",
					Name:     "John Doe",
					Username: "johndoe",
					Email:    "john@example.com",
					Password: "password123",
				}
			},
			wantError: false,
		},
		{
			name: "repository error",
			user: &domain_user.User{
				ID:       "test-user-id",
				Name:     "Updated Name",
				Username: "updateduser",
				Email:    "updated@example.com",
				Password: "newpassword",
			},
			setupMock: func(m *MockUserRepository) {
				m.updateError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository()
			tt.setupMock(mockRepo)
			mockBeltRepo := NewMockBeltProgressRepository()
			beltService := application_belt.NewBeltService(mockBeltRepo)
			service := application_user.NewUserService(mockRepo, *beltService)

			result, err := service.UpdateUser(tt.user)

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
				if result != nil && result.Name != tt.user.Name {
					t.Errorf("expected name %s but got %s", tt.user.Name, result.Name)
				}
			}
		})
	}
}

func TestUserService_AddBeltProgress(t *testing.T) {
	tests := []struct {
		name      string
		user      *domain_user.User
		request   *application_user.CreateUserRequest
		setupMock func(*MockUserRepository, *MockBeltProgressRepository)
		wantError bool
	}{
		{
			name: "successful belt progress addition",
			user: &domain_user.User{
				ID:   "test-user-id",
				Name: "John Doe",
			},
			request: &application_user.CreateUserRequest{
				BeltColor:  "blue",
				BeltStripe: 2,
			},
			setupMock: func(m *MockUserRepository, b *MockBeltProgressRepository) {},
			wantError: false,
		},
		{
			name: "belt service error",
			user: &domain_user.User{
				ID:   "test-user-id",
				Name: "John Doe",
			},
			request: &application_user.CreateUserRequest{
				BeltColor:  "blue",
				BeltStripe: 2,
			},
			setupMock: func(m *MockUserRepository, b *MockBeltProgressRepository) {
				b.createError = errors.New("belt service error")
			},
			wantError: true,
		},
		{
			name: "repository update error",
			user: &domain_user.User{
				ID:   "test-user-id",
				Name: "John Doe",
			},
			request: &application_user.CreateUserRequest{
				BeltColor:  "blue",
				BeltStripe: 2,
			},
			setupMock: func(m *MockUserRepository, b *MockBeltProgressRepository) {
				m.updateError = errors.New("database error")
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository()
			mockBeltRepo := NewMockBeltProgressRepository()
			tt.setupMock(mockRepo, mockBeltRepo)
			beltService := application_belt.NewBeltService(mockBeltRepo)
			service := application_user.NewUserService(mockRepo, *beltService)

			result, err := service.AddBeltProgress(tt.user, tt.request)

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
			}
		})
	}
}

