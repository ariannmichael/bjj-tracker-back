package application_user

import (
	"bjj-tracker/config"
	domain_user "bjj-tracker/src/modules/user/domain"
	infrastructure_user "bjj-tracker/src/modules/user/infrastructure"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserUseCase struct {
	Repo domain_user.UserRepository
}

func NewLoginUserUseCase() *LoginUserUseCase {
	db := config.ConnectToDB()
	repo := &infrastructure_user.UserRepositoryImpl{DB: db}

	return &LoginUserUseCase{
		Repo: repo,
	}
}

func (lc *LoginUserUseCase) Execute(req LoginUserRequest) (*string, error) {
	if req.Email == "" || req.Password == "" {
		return nil, fmt.Errorf("email, and password are required")
	}

	user, err := lc.Repo.FindByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find User by Email: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("wrong password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, fmt.Errorf("failed to create Token")
	}

	return &tokenString, nil
}
