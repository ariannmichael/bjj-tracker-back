package infrastructure_user

import (
	domain_user "bjj-tracker/src/modules/user/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

var _ domain_user.UserRepository = &UserRepositoryImpl{}

func (r *UserRepositoryImpl) Create(user *domain_user.User) (*domain_user.User, error) {
	var existingUser domain_user.User
	if err := r.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		panic("user already exists")
	}

	if err := r.DB.Create(user).Error; err != nil {
		panic(err)
	}
	return user, nil
}

// FindByEmail implements domain_user.UserRepository.
func (r *UserRepositoryImpl) FindByEmail(email string) (*domain_user.User, error) {
	panic("unimplemented")
}
