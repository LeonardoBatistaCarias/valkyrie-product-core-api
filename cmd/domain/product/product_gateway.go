package product

import (
	"context"
)

type ProductGateway interface {
	Create(ctx context.Context, product Product) error
	DeleteProductByID(ctx context.Context, product Product) error
}
