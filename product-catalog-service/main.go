package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/tung238/booking-microservice/product-catalog-service/application"
	"github.com/tung238/booking-microservice/product-catalog-service/infrastructure/db"
	"github.com/tung238/booking-microservice/product-catalog-service/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("product-catalog-service"),
	)
	service.Init()
	// collection := client.Database("booking").Collection("product_catalog")
	repo := application.NewRepository(db.DB)
	h := application.NewHandler(repo)
	proto.RegisterProductCatalogServiceHandler(service.Server(), h)

}
