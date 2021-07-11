package application

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	dm "github.com/tung238/booking-microservice/product-catalog-service/domain"
	proto "github.com/tung238/booking-microservice/product-catalog-service/proto"
)

type repository interface {
	ListProducts(ctx context.Context, req *proto.ListProductsRequest) (*dm.ListProductsResponse, error)
	GetProduct(ctx context.Context, req *proto.GetProductRequest) (*dm.Product, error)
	CreateProduct(ctx context.Context, req *proto.CreateProductRequest, res *proto.Product) (*dm.Product, error)
	UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest, res *proto.Product) (*dm.Product, error)
	DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) ListProducts(ctx context.Context, req *proto.ListProductsRequest) (*dm.ListProductsResponse, error) {
	return nil, nil
}

func (r *Repository) GetProduct(ctx context.Context, req *proto.GetProductRequest) (*dm.Product, error) {
	var p *dm.Product
	r.db.First(p, req.Id)
	if p == nil {
		return nil, errors.New("Product not found")
	}
	return p, nil
}

func (r *Repository) CreateProduct(ctx context.Context, req *proto.CreateProductRequest, res *proto.Product) (*dm.Product, error) {
	return nil, nil
}

func (r *Repository) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest, res *proto.Product) (*dm.Product, error) {
	return nil, nil
}

func (r *Repository) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) error {
	return nil
}
