package product

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Product struct {
	ProductID     uuid.UUID
	Name          string
	Description   string
	Brand         Brand
	Price         float64
	Quantity      int
	CategoryID    uuid.UUID
	ProductImages []*ProductImage
	Active        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

func NewProduct(
	productID uuid.UUID,
	name string,
	description string,
	brand Brand,
	price float64,
	quantity int,
	categoryID uuid.UUID,
	images []*ProductImage,
	active bool) *Product {
	return &Product{
		ProductID:     productID,
		Name:          name,
		Description:   description,
		Brand:         brand,
		Price:         price,
		Quantity:      quantity,
		CategoryID:    categoryID,
		ProductImages: images,
		Active:        active,
	}
}
