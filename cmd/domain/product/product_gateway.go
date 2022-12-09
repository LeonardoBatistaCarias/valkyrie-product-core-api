package product

import (
	"context"
)

type ProductGateway interface {
	CreateProduct(ctx context.Context, product Product) error
	DeleteProductByID(ctx context.Context, product Product) error
	UpdateProductByID(ctx context.Context, product Product) error
}
