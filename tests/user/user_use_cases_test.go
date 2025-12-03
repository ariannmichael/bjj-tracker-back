package user_test

import (
	"errors"
	"os"
	"testing"

	application_belt "bjj-tracker/src/modules/belt/application"
	application_user "bjj-tracker/src/modules/user/application"
	domain_user "bjj-tracker/src/modules/user/domain"

	"golang.org/x/crypto/bcrypt"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		request   application_user.CreateUserRequest
		setupMock func(*MockUserRepository, *MockBeltProgressRepository)
		wantError bool
	}{
		{
			name: "successful creation",
			request: application_user.CreateUserRequest{
				Name:       "John Doe",
				Username:   "johndoe",
				Email:      "john@example.com",
				Password:   "password123",
				BeltColor:  "blue",
				BeltStripe: 2,
			},
			setupMock: func(m *MockUserRepository, b *MockBeltProgressRepository) {},
			wantError: false,
		},
		{
			name: "missing required fields",
			request: application_user.CreateUserRequest{
				Name:       "",
				Email:      "",
				Password:   "",
				BeltColor:  "blue",
				BeltStripe: 2,
			},
			setupMock: func(m *MockUserRepository, b *MockBeltProgressRepository) {},
			wantError: true,
		},
		{
			name: "missing belt information",
			request: application_user.CreateUserRequest{
				Name:       "John Doe",
				Username:   "johndoe",
				Email:      "john@example.com",
				Password:   "password123",
				BeltColor:  "",
				BeltStripe: -1,
			},
			setupMock: func(m *MockUserRepository, b *MockBeltProgressRepository) {},
			wantError: true,
		},
		{
			name: "user service error",
			request: application_user.CreateUserRequest{
				Name:       "John Doe",
				Username:   "johndoe",
				Email:      "john@example.com",
				Password:   "password123",
				BeltColor:  "blue",
				BeltStripe: 2,
			},
			setupMock: func(m *MockUserRepository, b *MockBeltProgressRepository) {
				m.createError = errors.New("database error")
			},
			wantError: true,
		},
		{
			name: "belt progress error",
			request: application_user.CreateUserRequest{
				Name:       "John Doe",
				Username:   "johndoe",
				Email:      "john@example.com",
				Password:   "password123",
				BeltColor:  "blue",
				BeltStripe: 2,
			},
			setupMock: func(m *MockUserRepository, b *MockBeltProgressRepository) {
				b.createError = errors.New("belt service error")
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
			userService := application_user.NewUserService(mockRepo, *beltService)
			useCase := &application_user.CreateUserUseCase{
				Repo:        mockRepo,
				UserService: userService,
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

func TestLoginUserUseCase_Execute(t *testing.T) {
	// Set a test secret for JWT
	os.Setenv("SECRET", "test-secret-key-for-jwt-token-generation")

	tests := []struct {
		name      string
		request   application_user.LoginUserRequest
		setupMock func(*MockUserRepository)
		wantError bool
	}{
		{
			name: "successful login",
			request: application_user.LoginUserRequest{
				Email:    "john@example.com",
				Password: "password123",
			},
			setupMock: func(m *MockUserRepository) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
				m.users["test-user-id"] = &domain_user.User{
					ID:       "test-user-id",
					Email:    "john@example.com",
					Password: string(hashedPassword),
				}
			},
			wantError: false,
		},
		{
			name: "missing email and password",
			request: application_user.LoginUserRequest{
				Email:    "",
				Password: "",
			},
			setupMock: func(m *MockUserRepository) {},
			wantError: true,
		},
		{
			name: "user not found",
			request: application_user.LoginUserRequest{
				Email:    "notfound@example.com",
				Password: "password123",
			},
			setupMock: func(m *MockUserRepository) {},
			wantError: true,
		},
		{
			name: "wrong password",
			request: application_user.LoginUserRequest{
				Email:    "john@example.com",
				Password: "wrongpassword",
			},
			setupMock: func(m *MockUserRepository) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
				m.users["test-user-id"] = &domain_user.User{
					ID:       "test-user-id",
					Email:    "john@example.com",
					Password: string(hashedPassword),
				}
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository()
			tt.setupMock(mockRepo)
			useCase := &application_user.LoginUserUseCase{
				Repo: mockRepo,
			}

			result, err := useCase.Execute(tt.request)

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
					t.Errorf("expected token but got nil")
				}
				if result != nil && *result == "" {
					t.Errorf("expected non-empty token but got empty string")
				}
			}
		})
	}
}

func TestGetUserByIDUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		userID    string
		setupMock func(*MockUserRepository)
		wantError bool
	}{
		{
			name:   "successful retrieval",
			userID: "test-user-id",
			setupMock: func(m *MockUserRepository) {
				m.users["test-user-id"] = &domain_user.User{
					ID:       "test-user-id",
					Name:     "John Doe",
					Username: "johndoe",
					Email:    "john@example.com",
				}
			},
			wantError: false,
		},
		{
			name:      "user not found",
			userID:    "non-existent",
			setupMock: func(m *MockUserRepository) {},
			wantError: true,
		},
		{
			name:   "repository error",
			userID: "test-user-id",
			setupMock: func(m *MockUserRepository) {
				m.findByIDError = errors.New("database error")
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
			userService := application_user.NewUserService(mockRepo, *beltService)
			useCase := &application_user.GetUserByIDUseCase{
				Repo:        mockRepo,
				UserService: userService,
			}

			result, err := useCase.Execute(tt.userID)

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
				if result != nil && result.ID != tt.userID {
					t.Errorf("expected ID %s but got %s", tt.userID, result.ID)
				}
			}
		})
	}
}

func TestGetAllUsersUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(*MockUserRepository)
		wantCount int
		wantError bool
	}{
		{
			name: "successful retrieval",
			setupMock: func(m *MockUserRepository) {
				m.users["user-1"] = &domain_user.User{
					ID:       "user-1",
					Name:     "John Doe",
					Username: "johndoe",
					Email:    "john@example.com",
				}
				m.users["user-2"] = &domain_user.User{
					ID:       "user-2",
					Name:     "Jane Smith",
					Username: "janesmith",
					Email:    "jane@example.com",
				}
			},
			wantCount: 2,
			wantError: false,
		},
		{
			name:      "empty list",
			setupMock: func(m *MockUserRepository) {},
			wantCount: 0,
			wantError: false,
		},
		{
			name: "repository error",
			setupMock: func(m *MockUserRepository) {
				m.findAllError = errors.New("database error")
			},
			wantCount: 0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockUserRepository()
			tt.setupMock(mockRepo)
			mockBeltRepo := NewMockBeltProgressRepository()
			beltService := application_belt.NewBeltService(mockBeltRepo)
			userService := application_user.NewUserService(mockRepo, *beltService)
			useCase := &application_user.GetAllUsersUseCase{
				Repo:        mockRepo,
				UserService: userService,
			}

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
					t.Errorf("expected %d users but got %d", tt.wantCount, len(result))
				}
			}
		})
	}
}

func TestUpdateUserByIDUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		userID    string
		request   application_user.UpdateUserByIDRequest
		setupMock func(*MockUserRepository)
		wantError bool
	}{
		{
			name:   "successful update",
			userID: "test-user-id",
			request: application_user.UpdateUserByIDRequest{
				Name:       "Updated Name",
				Username:   "updateduser",
				Email:      "updated@example.com",
				Password:   "newpassword123",
				BeltColor:  "purple",
				BeltStripe: 3,
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
			name:   "user not found",
			userID: "non-existent",
			request: application_user.UpdateUserByIDRequest{
				Name:       "Updated Name",
				Username:   "updateduser",
				Email:      "updated@example.com",
				Password:   "newpassword123",
				BeltColor:  "purple",
				BeltStripe: 3,
			},
			setupMock: func(m *MockUserRepository) {},
			wantError: true,
		},
		{
			name:   "repository find error",
			userID: "test-user-id",
			request: application_user.UpdateUserByIDRequest{
				Name:       "Updated Name",
				Username:   "updateduser",
				Email:      "updated@example.com",
				Password:   "newpassword123",
				BeltColor:  "purple",
				BeltStripe: 3,
			},
			setupMock: func(m *MockUserRepository) {
				m.findByIDError = errors.New("database error")
			},
			wantError: true,
		},
		{
			name:   "repository update error",
			userID: "test-user-id",
			request: application_user.UpdateUserByIDRequest{
				Name:       "Updated Name",
				Username:   "updateduser",
				Email:      "updated@example.com",
				Password:   "newpassword123",
				BeltColor:  "purple",
				BeltStripe: 3,
			},
			setupMock: func(m *MockUserRepository) {
				m.users["test-user-id"] = &domain_user.User{
					ID:       "test-user-id",
					Name:     "John Doe",
					Username: "johndoe",
					Email:    "john@example.com",
					Password: "password123",
				}
				m.updateError = errors.New("database update error")
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
			userService := application_user.NewUserService(mockRepo, *beltService)
			useCase := &application_user.UpdateUserByIDUseCase{
				Repo:        mockRepo,
				UserService: userService,
			}

			result, err := useCase.Execute(tt.userID, tt.request)

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
				if result != nil {
					if result.Name != tt.request.Name {
						t.Errorf("expected name %s but got %s", tt.request.Name, result.Name)
					}
					if result.Email != tt.request.Email {
						t.Errorf("expected email %s but got %s", tt.request.Email, result.Email)
					}
				}
			}
		})
	}
}
