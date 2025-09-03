package infrastructure_user

import (
	domain_user "bjj-tracker/src/modules/user/domain"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

var _ domain_user.UserRepository = &UserRepositoryImpl{}

func (r *UserRepositoryImpl) Create(user *domain_user.User) (*domain_user.User, error) {
	var existingUser domain_user.User
	if err := r.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("user with email %s already exists", user.Email)
	}

	// Generate UUID for new user
	user.ID = uuid.New().String()

	if err := r.DB.Create(user).Error; err != nil {
		return nil, fmt.Errorf("REPO failed to create user: %w", err)
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*domain_user.User, error) {
	var user domain_user.User
	if err := r.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, fmt.Errorf("REPO failed to find user: %w", err)
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Update(user *domain_user.User) (*domain_user.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByID(id string) (*domain_user.User, error) {
	var user domain_user.User
	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("REPO failed to find user: %w", err)
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindAll() ([]domain_user.User, error) {
	var users []domain_user.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("REPO failed to find all users: %w", err)
	}
	return users, nil
}
