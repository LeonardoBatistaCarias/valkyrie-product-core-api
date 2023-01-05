package category

import (
	"context"
)

type CategoryGateway interface {
	FindCategoryByID(ctx context.Context, categoryID string) (*Category, error)
}
