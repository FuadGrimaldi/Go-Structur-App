package service

import (
	"context"
	"go-app/internal/dto"
	"go-app/internal/entity"
	"go-app/internal/repository"
)

type ProductService interface {
	FindAll(ctx context.Context) ([]dto.Product, error)
	Create(ctx context.Context, req dto.NewProduct) error
	FindOneById(ctx context.Context, id int64) (*dto.Product, error)
}

type productService struct {
	repository	repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &productService{repository}
}

func (ps *productService) FindAll(ctx context.Context) ([]dto.Product, error) {
	products, err := ps.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	productDto := make([]dto.Product, 0)
	for _, v := range products {
		productDto = append(productDto, dto.Product{ID: v.ID, Title: v.Title, Author: v.Author, Publicatio_year: v.Publicatio_year, Description: v.Description, Category: v.Category, ISBN: v.ISBN, Stoct: v.Stoct, Price: v.Price })
	}
	return productDto, nil
}

func (ps *productService) FindOneById(ctx context.Context, id int64) (*dto.Product, error) {
	product, err := ps.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	productDto := &dto.Product{ID: product.ID, Title: product.Title, Author: product.Author, Publicatio_year: product.Publicatio_year, Description: product.Description, Category: product.Category, ISBN: product.ISBN, Stoct: product.Stoct, Price: product.Price }
	return productDto, nil
}

func (ps *productService) Create(ctx context.Context, req dto.NewProduct) error {
	product := entity.Product{
		Title: req.Title,
		Author: req.Author,
		Publicatio_year: req.Publicatio_year,
		Description: req.Description,
		Category: req.Category,
		ISBN: req.ISBN,
		Stoct: req.Stoct,
		Price: req.Price,
	}
	return ps.repository.Create(ctx, &product)
}