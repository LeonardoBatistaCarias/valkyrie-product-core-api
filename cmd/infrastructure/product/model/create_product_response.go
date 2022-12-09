package model

import uuid "github.com/satori/go.uuid"

type CreateProductResponse struct {
	ProductID uuid.UUID `json:"product_id"`
}

func NewCreateProductResponse(productID uuid.UUID) *CreateProductResponse {
	return &CreateProductResponse{ProductID: productID}
}
