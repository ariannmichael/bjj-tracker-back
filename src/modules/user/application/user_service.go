package application_user

import (
	application_belt "bjj-tracker/src/modules/belt/application"
	domain_user "bjj-tracker/src/modules/user/domain"
	"fmt"
)

type UserService struct {
	Repo        domain_user.UserRepository
	BeltService application_belt.BeltService
}

func NewUserService(repo domain_user.UserRepository, beltService application_belt.BeltService) *UserService {
	return &UserService{
		Repo:        repo,
		BeltService: beltService,
	}
}

func (us *UserService) CreateUser(userDTO *CreateUserRequest) (*domain_user.User, error) {
	user := domain_user.User{
		Name:     userDTO.Name,
		Username: userDTO.Username,
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}

	newUser, err := us.Repo.Create(&user)
	if err != nil {
		return nil, fmt.Errorf("SERVICE failed to create user: %w", err)
	}

	return newUser, nil
}

func (us *UserService) AddBeltProgress(user *domain_user.User, userDTO *CreateUserRequest) (*domain_user.User, error) {
	createBeltProgressDTO := application_belt.CreateBeltProgressDTO{
		UserID:  user.ID,
		Color:   userDTO.BeltColor,
		Stripes: userDTO.BeltStripe,
	}
	beltProgress, err := us.BeltService.CreateBeltProgress(createBeltProgressDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to create belt progress: %w", err)
	}
	user.BeltProgress = append(user.BeltProgress, *beltProgress)
	updatedUser, err := us.Repo.Update(user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user with belt progress: %w", err)
	}

	return updatedUser, nil
}
