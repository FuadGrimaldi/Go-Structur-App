package repository

import (
	"context"
	"go-app/internal/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
	Create(ctx context.Context, product *entity.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (pr *productRepository) FindAll(ctx context.Context) ([]entity.Product, error) {
	product := make([]entity.Product, 0)

	if err := pr.db.WithContext(ctx).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (pr *productRepository) Create(ctx context.Context, product *entity.Product) error {
	if err := pr.db.WithContext(ctx).Create(&product).Error; err != nil {
		return err
	}
	return nil
}