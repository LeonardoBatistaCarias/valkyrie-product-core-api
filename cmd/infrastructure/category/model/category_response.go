package model

import "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/category"

type CategoryResponse struct {
	ID   string
	Name string
}

func (c *CategoryResponse) ToModel() *category.Category {
	return &category.Category{
		ID:   c.ID,
		Name: c.Name,
	}
}
