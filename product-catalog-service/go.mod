module github.com/tung238/booking-microservice/product-catalog-service

go 1.16

require (
	github.com/asim/go-micro/v3 v3.5.2
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/configor v1.2.1
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.mongodb.org/mongo-driver v1.5.4 // indirect
	google.golang.org/protobuf v1.27.1
)

replace github.com/tung238/booking-microservice/product-catalog-service => ../product-catalog-service
