package service

import (
	"context"
	"errors"
	"go-app/config"
	"go-app/internal/common"
	"go-app/internal/dto"
	"go-app/internal/repository"

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

	token, err := common.GenerateAccessToken(ctx, user)
	if err != nil {
		return "", err
	}
	return token, nil
}