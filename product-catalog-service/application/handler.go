package application

import (
	"context"

	proto "github.com/tung238/booking-microservice/product-catalog-service/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	r repository
}

func NewHandler(r repository) *Handler {
	return &Handler{r: r}
}

func (h *Handler) ListProducts(ctx context.Context, req *proto.ListProductsRequest, res *proto.ListProductsResponse) error {
	// resp, err := h.readModel.ListProducts(ctx, req)
	// res = resp
	return nil
}

func (h *Handler) GetProduct(ctx context.Context, req *proto.GetProductRequest, res *proto.Product) error {
	_, err := h.r.GetProduct(ctx, req)

	return err
}

func (h *Handler) CreateProduct(ctx context.Context, req *proto.CreateProductRequest, res *proto.Product) error {
	return nil

}

func (h *Handler) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest, res *proto.Product) error {
	return nil
}

func (h *Handler) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest, res *emptypb.Empty) error {
	return nil
}
