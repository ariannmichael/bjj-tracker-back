package application_user

import (
	"fmt"

	domain_user "bjj-tracker/src/modules/user/domain"
)

type CreateUserUseCase struct {
	Repo domain_user.UserRepository
}

func (uc *CreateUserUseCase) Execute(name string, email string, password string) (*domain_user.User, error) {
	if name == "" || email == "" || password == "" {
		return nil, fmt.Errorf("name, email, and password are required")
	}

	user := &domain_user.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	return uc.Repo.Create(user)
}
