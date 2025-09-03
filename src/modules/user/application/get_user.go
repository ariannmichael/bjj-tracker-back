package application_user

import (
	"bjj-tracker/config"
	application_belt "bjj-tracker/src/modules/belt/application"
	infrastructure_belt "bjj-tracker/src/modules/belt/infrastructure"
	domain_user "bjj-tracker/src/modules/user/domain"
	infrastructure_user "bjj-tracker/src/modules/user/infrastructure"
	"fmt"
)

type GetUserByIDUseCase struct {
	Repo        domain_user.UserRepository
	UserService *UserService
}

func NewGetUserByIDUseCase() *GetUserByIDUseCase {
	db := config.ConnectToDB()
	repo := &infrastructure_user.UserRepositoryImpl{DB: db}
	beltRepo := infrastructure_belt.NewBeltProgressRepository(db)
	beltService := application_belt.NewBeltService(beltRepo)
	userService := NewUserService(repo, *beltService)
	return &GetUserByIDUseCase{
		Repo:        repo,
		UserService: userService,
	}
}

func (uc *GetUserByIDUseCase) Execute(userID string) (*domain_user.User, error) {
	user, err := uc.Repo.FindByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return user, nil
}
