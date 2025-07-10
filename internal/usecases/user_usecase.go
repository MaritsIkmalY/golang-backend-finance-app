package usecases

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maritsikmaly/golang-finance-app/internal/entities"
	"github.com/maritsikmaly/golang-finance-app/internal/models"
	"github.com/maritsikmaly/golang-finance-app/internal/models/converter"
	"github.com/maritsikmaly/golang-finance-app/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(req *models.RegisterUserRequest) (*models.UserResponse, error)
	Login(req *models.LoginUserRequest) (*models.UserResponse, error)
}

type userUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(ur repositories.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: ur,
	}
}

func (u *userUseCase) Register(req *models.RegisterUserRequest) (*models.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hashedPassword),
	}

	if err := u.userRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := u.generateToken(user.ID)

	if err != nil {
		return nil, err
	}

	return converter.UserTokenResponse(user, token), nil
}

func (u *userUseCase) Login(req *models.LoginUserRequest) (*models.UserResponse, error) {
	userResponse, err := u.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userResponse.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := u.generateToken(userResponse.ID)
	if err != nil {
		return nil, err
	}

	return converter.UserTokenResponse(userResponse, token), nil
}

func (u *userUseCase) generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(secret))
}
