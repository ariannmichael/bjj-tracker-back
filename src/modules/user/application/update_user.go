package application_user

import (
	"bjj-tracker-monorepo/bjj-tracker/config"
	infrastructure_belt "bjj-tracker-monorepo/bjj-tracker/src/modules/belt/infrastructure"
	infrastructure_user "bjj-tracker-monorepo/bjj-tracker/src/modules/user/infrastructure"
	application_belt "bjj-tracker/src/modules/belt/application"
	domain_user "bjj-tracker/src/modules/user/domain"
	"fmt"
)

type UpdateUserByIDUseCase struct {
	Repo        domain_user.UserRepository
	UserService *UserService
}

func NewUpdateUserByIDUseCase() *UpdateUserByIDUseCase {
	db := config.ConnectToDB()
	repo := &infrastructure_user.UserRepositoryImpl{DB: db}
	beltRepo := infrastructure_belt.NewBeltProgressRepository(db)
	beltService := application_belt.NewBeltService(beltRepo)
	userService := NewUserService(repo, *beltService)
	return &UpdateUserByIDUseCase{
		Repo:        repo,
		UserService: userService,
	}
}

func (uc *UpdateUserByIDUseCase) Execute(userID string, req UpdateUserByIDRequest) (*domain_user.User, error) {
	user, err := uc.Repo.FindByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	user.Name = req.Name
	user.Username = req.Username
	user.Email = req.Email
	user.Password = req.Password
	user.BeltColor = req.BeltColor
	user.BeltStripe = req.BeltStripe

	newUser, err := uc.UserService.UpdateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	if newUser == nil {
		return nil, fmt.Errorf("failed to update user: user is nil")
	}
	return newUser, nil
}
