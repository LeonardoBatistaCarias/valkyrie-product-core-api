package model

import uuid "github.com/satori/go.uuid"

type UpdateProductResponse struct {
	ProductID uuid.UUID `json:"product_id"`
}

func NewUpdateProductResponse(productID uuid.UUID) *UpdateProductResponse {
	return &UpdateProductResponse{ProductID: productID}
}
