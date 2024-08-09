package service

import (
	"context"
	"errors"
	"fmt"
	"go-app/config"
	"go-app/internal/common"
	"go-app/internal/dto"
	"go-app/internal/entity"
	"go-app/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, request dto.LoginRequest) (string, error)
	FindAll(ctx context.Context) ([]dto.User, error)
	FindOne(ctx context.Context, id int64) (*dto.User, error)
	Create(ctx context.Context, request dto.NewUser) error
	Update(ctx context.Context, request dto.UpdateUser) error
	Delete(ctx context.Context, id int64) error
}

type userService struct {
	cfg        *config.Config
	repository repository.UserRepository
}

func NewUserService(cfg *config.Config, repository repository.UserRepository) UserService {
	return &userService{cfg, repository}
}

func (u *userService) Login(ctx context.Context, request dto.LoginRequest) (string, error) {
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

func (u *userService) FindAll(ctx context.Context) ([]dto.User, error) {
	users, err := u.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	usersDTO := make([]dto.User, 0)
	for _, v := range users {
		usersDTO = append(usersDTO, dto.User{ID: v.ID, Name: v.Name, Address: v.Address, Gender: v.Gender, Email: v.Email, Username: v.Username})
	}
	return usersDTO, nil
}

func (u *userService) FindOne(ctx context.Context, id int64) (*dto.User, error) {
	user, err := u.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	userDTO := &dto.User{ID: user.ID, Name: user.Name, Address: user.Address, Gender: user.Gender, Email: user.Email, Username: user.Username}
	return userDTO, nil
}

func (u *userService) Create(ctx context.Context, request dto.NewUser)error{
	user := entity.User{
		Name: request.Name,
		Address: request.Address,
		Gender: request.Gender,
		Email: request.Email,
		Username: request.Username,
		Password: request.Password,
	}
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(password)
	
	// Create user using repository
	return u.repository.Create(ctx, &user)
}

func (u *userService) Update(ctx context.Context, request dto.UpdateUser)error {
	user, err := u.repository.FindByID(ctx, request.ID)
	fmt.Println(request.ID)
	if err != nil {
		return err
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Username != "" {
		user.Username = request.Username
	}

	if request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user.Password = string(password)
	}

	return u.repository.Update(ctx, user)
}

func (u *userService) Delete(ctx context.Context, id int64) error {
	return u.repository.Delete(ctx, id)
}