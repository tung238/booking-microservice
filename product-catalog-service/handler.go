package main

import (
	"context"

	proto "github.com/tung238/booking-service/product-catalog-service/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type handler struct {
	repository
	readModel
}

func (h *handler) ListProducts(ctx context.Context, req *proto.ListProductsRequest, res *proto.ListProductsResponse) error {
	// resp, err := h.readModel.ListProducts(ctx, req)
	// res = resp
	return nil
}

func (h *handler) GetProduct(ctx context.Context, req *proto.GetProductRequest, res *proto.Product) error {
	product, err := h.repository.GetProduct(ctx, req)
	res = product
	return err
}

func (h *handler) CreateProduct(ctx context.Context, req *proto.CreateProductRequest, res *proto.Product) error {
	return nil

}

func (h *handler) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest, res *proto.Product) error {
	return nil
}

func (h *handler) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest, res *emptypb.Empty) error {
	return nil
}
