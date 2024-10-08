package service

import (
	"context"
	"go-app/internal/dto"
	"go-app/internal/entity"
	"go-app/internal/repository"
	"strings"
	"time"
)

type ProductService interface {
	FindAll(ctx context.Context) ([]dto.Product, error)
	Create(ctx context.Context, req dto.NewProduct) error
	FindOneById(ctx context.Context, id int64) (*dto.Product, error)
	FindOneByTitle(ctx context.Context, title string) (*dto.Product, error)
	Update(ctx context.Context, req dto.UpdateProduct) error
	Delete(ctx context.Context, id int64) error
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

func (ps *productService) FindOneByTitle(ctx context.Context, title string) (*dto.Product, error) {
	titleStr := strings.ReplaceAll(title, "-", " ")
	product, err := ps.repository.FindByTitle(ctx, titleStr)
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

func (ps *productService) Update(ctx context.Context, req dto.UpdateProduct) error {
	product, err := ps.repository.FindById(ctx, req.ID)
	if err != nil {
		return err
	}

	if req.Title != "" {product.Title = req.Title}
	if req.Author != "" {product.Author = req.Author}
	if req.Publicatio_year != 0 {product.Publicatio_year = req.Publicatio_year}
	if req.Description != "" {product.Description = req.Description}
	if req.Category != "" {product.Category = req.Category}
	if req.ISBN != "" {product.ISBN = req.ISBN}
	if req.Stoct != 0 {product.Stoct = req.Stoct}
	if req.Price != 0 {product.Price = req.Price}

	product.UpdatedAt = time.Now()

	return ps.repository.Update(ctx, product)
}

func (ps *productService) Delete(ctx context.Context, id int64) error {
	return ps.repository.Delete(ctx, id)
}