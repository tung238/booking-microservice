package main

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	proto "github.com/tung238/booking-service/product-catalog-service/proto"
)

type repository interface {
	ListProducts(ctx context.Context, req *proto.ListProductsRequest) (ListProductsResponse, error)
	GetProduct(ctx context.Context, req *proto.GetProductRequest) (*Product, error)
	CreateProduct(ctx context.Context, req *proto.CreateProductRequest, res *proto.Product) (Product, error)
	UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest, res *proto.Product) (Product, error)
	DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) error
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) ListProducts(ctx context.Context, req *proto.ListProductsRequest) (*ListProductsResponse, error) {
	return nil, nil
}

func (r *Repository) GetProduct(ctx context.Context, req *proto.GetProductRequest) (*Product, error) {
	var p *Product
	r.DB.First(p, req.id)
	if p == nil {
		return nil, errors.New("Product not found")
	}
	return p, nil
}

func (r *Repository) CreateProduct(ctx context.Context, req *proto.CreateProductRequest, res *proto.Product) (*Product, error) {
	return nil, nil
}

func (r *Repository) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest, res *proto.Product) (*Product, error) {
	return nil, nil
}

func (r *Repository) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) error {
	return nil
}
