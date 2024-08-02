package repository

import (
	"go-app/internal/entity"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	users := make([]entity.User, 0)

	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

