package application_user

import (
	"fmt"

	"bjj-tracker/config"
	application_belt "bjj-tracker/src/modules/belt/application"
	infrastructure_belt "bjj-tracker/src/modules/belt/infrastructure"
	domain_user "bjj-tracker/src/modules/user/domain"
	infrastructure_user "bjj-tracker/src/modules/user/infrastructure"
)

type CreateUserUseCase struct {
	Repo        domain_user.UserRepository
	UserService *UserService
}

func NewCreateUserUseCase() *CreateUserUseCase {
	db := config.ConnectToDB()
	repo := &infrastructure_user.UserRepositoryImpl{DB: db}
	beltRepo := infrastructure_belt.NewBeltProgressRepository(db)
	beltService := application_belt.NewBeltService(beltRepo)
	userService := NewUserService(repo, *beltService)

	return &CreateUserUseCase{
		Repo:        repo,
		UserService: userService,
	}
}

func (uc *CreateUserUseCase) Execute(req CreateUserRequest) (*domain_user.User, error) {
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return nil, fmt.Errorf("name, email, and password are required")
	}

	if req.BeltColor == "" || req.BeltStripe < 0 {
		return nil, fmt.Errorf("belt color and stripe are required")
	}

	user, err := uc.UserService.CreateUser(&req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	newUser, err := uc.UserService.AddBeltProgress(user, &req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if newUser == nil {
		return nil, fmt.Errorf("failed to add belt progress: user is nil")
	}

	return newUser, nil
}
