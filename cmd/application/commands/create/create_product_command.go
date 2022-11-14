package create

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreateProductCommand struct {
	Name          string
	Description   string
	Brand         int32
	Price         float32
	Quantity      int32
	CategoryID    uuid.UUID
	ProductImages []*CreateProductImageCommand
	Active        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type CreateProductImageCommand struct {
	Address string
}

func NewCreateProductCommand(
	name string,
	description string,
	brand int32,
	price float32,
	quantity int32,
	categoryID uuid.UUID,
	images []*CreateProductImageCommand,
	active bool) *CreateProductCommand {
	return &CreateProductCommand{
		Name:          name,
		Description:   description,
		Brand:         brand,
		Price:         price,
		Quantity:      quantity,
		CategoryID:    categoryID,
		ProductImages: images,
		Active:        active,
	}
}
