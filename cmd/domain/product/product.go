package product

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Product struct {
	ProductID     uuid.UUID `json:"product_id" validate:"required"`
	Name          string    `json:"name" validate:"required,gte=5,lte=50"`
	Description   string    `json:"description" validate:"required,gte=20,lte=400"`
	Brand         Brand     `json:"brand" validate:"required"`
	Price         float32   `json:"price" validate:"required,gte=0"`
	Quantity      int32     `json:"quantity" validate:"required,gte=0"`
	CategoryID    uuid.UUID `json:"category_id" validate:"required,gte=0"`
	ProductImages []*Image  `json:"product_images" validate:"required,gte=0"`
	Active        bool      `json:"active" validate:"required"`
	CreatedAt     time.Time `json:"created_at" validate:"required,datetime"`
	UpdatedAt     time.Time `json:"updated_at" validate:"required,datetime"`
	DeletedAt     time.Time `json:"deleted_at" validate:"required,datetime"`
}

func NewProduct(
	productID uuid.UUID,
	name string,
	description string,
	brand Brand,
	price float32,
	quantity int32,
	categoryID uuid.UUID,
	images []*Image,
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
