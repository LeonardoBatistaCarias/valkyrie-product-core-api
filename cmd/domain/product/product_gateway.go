package product

import (
	"context"
	uuid "github.com/satori/go.uuid"
)

type ProductGateway interface {
	CreateProduct(ctx context.Context, product Product) error
	DeleteProductByID(ctx context.Context, productID uuid.UUID) error
	DeactivateProductByID(ctx context.Context, productID uuid.UUID) error
	UpdateProductByID(ctx context.Context, product Product) error
}
