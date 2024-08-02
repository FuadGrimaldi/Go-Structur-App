package service

import (
	"context"
	"go-app/internal/dto"
	"go-app/internal/repository"
)

type UserService interface {
	FindAll(ctx context.Context) ([]dto.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository}
}

func (u *userService) FindAll(ctx context.Context) ([]dto.User, error) {
	users, err := u.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	usersDTO := make([]dto.User, 0)
	for _, v := range users {
		usersDTO = append(usersDTO, dto.User{ID: v.ID, Name: v.Name, Address: v.Address, Gender: v.Gender, Email: v.Email})
	}
	return usersDTO, nil
}