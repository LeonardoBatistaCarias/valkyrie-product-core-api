package models

import (
	uuid "github.com/satori/go.uuid"
)

type CreateProductDTO struct {
	ProductID   uuid.UUID
	Name        string
	Description string
	Brand       int32
	Price       float32
	Quantity    int32
	CategoryID  uuid.UUID
}

type CreateProductResponseDTO struct {
	ProductID uuid.UUID
}
