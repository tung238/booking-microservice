package domain

import "github.com/jinzhu/gorm"

type ProductType string

const (
	SERVICE ProductType = "Service"
	PRODUCT ProductType = "Product"
)

type Product struct {
	gorm.Model
	Name         string
	NameWithSlug string
	Feature      bool
	Code         string
	CategoryId   int32
	MadeCountry  string
	Gender       string
	MainImage    string
	Price        float32
	Description  string
	Type         int16
}

type ListProductsResponse struct {
	TotalSize     int32
	NextPageToken int32
	Items         []Product
}
