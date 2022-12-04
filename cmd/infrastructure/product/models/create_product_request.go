package models

import (
	uuid "github.com/satori/go.uuid"
)

type CreateProductRequest struct {
	ProductID   uuid.UUID
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Brand       int32     `json:"brand"`
	Price       float32   `json:"price"`
	Quantity    int32     `json:"quantity"`
	CategoryID  uuid.UUID `json:"category_id"`
	Active      bool      `json:"active"`
	Images      [][]byte
}
