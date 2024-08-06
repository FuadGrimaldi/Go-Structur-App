package repository

import (
	"go-app/internal/entity"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	// Update(ctx context.Context) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	// Delete(ctx context.Context, id int64) (*entity.User, error)
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

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := new(entity.User)

	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

