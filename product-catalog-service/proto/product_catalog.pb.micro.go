// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: product_catalog.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ProductCatalogService service

func NewProductCatalogServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ProductCatalogService service

type ProductCatalogService interface {
	ListProducts(ctx context.Context, in *ListProductsRequest, opts ...client.CallOption) (*ListProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...client.CallOption) (*Product, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...client.CallOption) (*Product, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...client.CallOption) (*Product, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type productCatalogService struct {
	c    client.Client
	name string
}

func NewProductCatalogService(name string, c client.Client) ProductCatalogService {
	return &productCatalogService{
		c:    c,
		name: name,
	}
}

func (c *productCatalogService) ListProducts(ctx context.Context, in *ListProductsRequest, opts ...client.CallOption) (*ListProductsResponse, error) {
	req := c.c.NewRequest(c.name, "ProductCatalogService.ListProducts", in)
	out := new(ListProductsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogService) GetProduct(ctx context.Context, in *GetProductRequest, opts ...client.CallOption) (*Product, error) {
	req := c.c.NewRequest(c.name, "ProductCatalogService.GetProduct", in)
	out := new(Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogService) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...client.CallOption) (*Product, error) {
	req := c.c.NewRequest(c.name, "ProductCatalogService.CreateProduct", in)
	out := new(Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogService) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...client.CallOption) (*Product, error) {
	req := c.c.NewRequest(c.name, "ProductCatalogService.UpdateProduct", in)
	out := new(Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogService) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "ProductCatalogService.DeleteProduct", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductCatalogService service

type ProductCatalogServiceHandler interface {
	ListProducts(context.Context, *ListProductsRequest, *ListProductsResponse) error
	GetProduct(context.Context, *GetProductRequest, *Product) error
	CreateProduct(context.Context, *CreateProductRequest, *Product) error
	UpdateProduct(context.Context, *UpdateProductRequest, *Product) error
	DeleteProduct(context.Context, *DeleteProductRequest, *emptypb.Empty) error
}

func RegisterProductCatalogServiceHandler(s server.Server, hdlr ProductCatalogServiceHandler, opts ...server.HandlerOption) error {
	type productCatalogService interface {
		ListProducts(ctx context.Context, in *ListProductsRequest, out *ListProductsResponse) error
		GetProduct(ctx context.Context, in *GetProductRequest, out *Product) error
		CreateProduct(ctx context.Context, in *CreateProductRequest, out *Product) error
		UpdateProduct(ctx context.Context, in *UpdateProductRequest, out *Product) error
		DeleteProduct(ctx context.Context, in *DeleteProductRequest, out *emptypb.Empty) error
	}
	type ProductCatalogService struct {
		productCatalogService
	}
	h := &productCatalogServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductCatalogService{h}, opts...))
}

type productCatalogServiceHandler struct {
	ProductCatalogServiceHandler
}

func (h *productCatalogServiceHandler) ListProducts(ctx context.Context, in *ListProductsRequest, out *ListProductsResponse) error {
	return h.ProductCatalogServiceHandler.ListProducts(ctx, in, out)
}

func (h *productCatalogServiceHandler) GetProduct(ctx context.Context, in *GetProductRequest, out *Product) error {
	return h.ProductCatalogServiceHandler.GetProduct(ctx, in, out)
}

func (h *productCatalogServiceHandler) CreateProduct(ctx context.Context, in *CreateProductRequest, out *Product) error {
	return h.ProductCatalogServiceHandler.CreateProduct(ctx, in, out)
}

func (h *productCatalogServiceHandler) UpdateProduct(ctx context.Context, in *UpdateProductRequest, out *Product) error {
	return h.ProductCatalogServiceHandler.UpdateProduct(ctx, in, out)
}

func (h *productCatalogServiceHandler) DeleteProduct(ctx context.Context, in *DeleteProductRequest, out *emptypb.Empty) error {
	return h.ProductCatalogServiceHandler.DeleteProduct(ctx, in, out)
}
