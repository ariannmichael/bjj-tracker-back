package application_user

import (
	"bjj-tracker/config"
	application_belt "bjj-tracker/src/modules/belt/application"
	infrastructure_belt "bjj-tracker/src/modules/belt/infrastructure"
	domain_user "bjj-tracker/src/modules/user/domain"
	infrastructure_user "bjj-tracker/src/modules/user/infrastructure"
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

	newUser, err := uc.UserService.UpdateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	if newUser == nil {
		return nil, fmt.Errorf("failed to update user: user is nil")
	}

	// Update belt progress if provided
	if req.BeltColor != "" && req.BeltStripe >= 0 {
		createBeltProgressDTO := application_belt.CreateBeltProgressDTO{
			UserID:  newUser.ID,
			Color:   req.BeltColor,
			Stripes: req.BeltStripe,
		}
		beltProgress, err := uc.UserService.BeltService.CreateBeltProgress(createBeltProgressDTO)
		if err != nil {
			return nil, fmt.Errorf("failed to create belt progress: %w", err)
		}
		newUser.BeltProgress = append(newUser.BeltProgress, *beltProgress)
		updatedUser, err := uc.Repo.Update(newUser)
		if err != nil {
			return nil, fmt.Errorf("failed to update user with belt progress: %w", err)
		}
		return updatedUser, nil
	}

	return newUser, nil
}
