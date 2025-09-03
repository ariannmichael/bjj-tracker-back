package application_user

import (
	"bjj-tracker/config"
	application_belt "bjj-tracker/src/modules/belt/application"
	infrastructure_belt "bjj-tracker/src/modules/belt/infrastructure"
	domain_user "bjj-tracker/src/modules/user/domain"
	infrastructure_user "bjj-tracker/src/modules/user/infrastructure"
	"fmt"
)

type GetAllUsersUseCase struct {
	Repo        domain_user.UserRepository
	UserService *UserService
}

func NewGetAllUsersUseCase() *GetAllUsersUseCase {
	db := config.ConnectToDB()
	repo := &infrastructure_user.UserRepositoryImpl{DB: db}
	beltRepo := infrastructure_belt.NewBeltProgressRepository(db)
	beltService := application_belt.NewBeltService(beltRepo)
	userService := NewUserService(repo, *beltService)

	return &GetAllUsersUseCase{
		Repo:        repo,
		UserService: userService,
	}
}

func (uc *GetAllUsersUseCase) Execute() ([]domain_user.User, error) {
	users, err := uc.Repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all users: %w", err)
	}
	return users, nil
}
