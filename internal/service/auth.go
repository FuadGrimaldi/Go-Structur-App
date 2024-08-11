package service

import (
	"context"
	"errors"
	"go-app/config"
	"go-app/internal/common"
	"go-app/internal/dto"
	"go-app/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, request dto.LoginRequest) (string, error)
}

type authService struct {
	cfg        *config.Config
	repository repository.UserRepository
}

func NewAuthService(cfg *config.Config, repository repository.UserRepository) AuthService {
	return &authService{cfg, repository}
}

func (u *authService) Login(ctx context.Context, request dto.LoginRequest) (string, error) {
	user, err := u.repository.FindByUsername(ctx, request.Username)

	if err != nil {
		return "", errors.New("username/password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return "", errors.New("username/password salah")
	}

	expiredTime := time.Now().Local().Add(10 * time.Minute)
	claims := common.JwtCustomClaims{
		Username: user.Username,
		Name:     user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := token.SignedString([]byte(u.cfg.JWTSecretKey))

	if err != nil {
		return "", err
	}

	return encodedToken, nil
}