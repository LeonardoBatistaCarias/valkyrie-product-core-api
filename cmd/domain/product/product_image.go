package product

import uuid "github.com/satori/go.uuid"

type Image struct {
	Name      string    `json:"name" validate:"required,gte=5,lte=20"`
	ProductID uuid.UUID `json:"product_id" validate:"required"`
}

func NewImage(name string, productId uuid.UUID) *Image {
	newFileName := uuid.NewV5(productId, name)
	return &Image{Name: newFileName.String(), ProductID: productId}
}
