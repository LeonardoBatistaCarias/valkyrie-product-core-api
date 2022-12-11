package model

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/model"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProductCommand struct {
	ProductID   uuid.UUID
	Name        string
	Description string
	Brand       int32
	Price       float32
	Quantity    int32
	CategoryID  uuid.UUID
	Images      []*ImageCommand
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewProductCommandToCreate(req model.CreateProductRequest) *ProductCommand {
	return NewProductCommand(
		req.ProductID,
		req.Name,
		req.Description,
		req.Brand,
		req.Price,
		req.Quantity,
		req.CategoryID,
		nil,
		req.Active,
	)
}

func NewProductCommandToUpdate(req model.UpdateProductRequest) *ProductCommand {
	return NewProductCommand(
		req.ProductID,
		req.Name,
		req.Description,
		req.Brand,
		req.Price,
		req.Quantity,
		req.CategoryID,
		nil,
		req.Active,
	)
}

func NewProductCommand(
	productID uuid.UUID,
	name string,
	description string,
	brand int32,
	price float32,
	quantity int32,
	categoryID uuid.UUID,
	images []*ImageCommand,
	active bool) *ProductCommand {
	return &ProductCommand{
		ProductID:   productID,
		Name:        name,
		Description: description,
		Brand:       brand,
		Price:       price,
		Quantity:    quantity,
		CategoryID:  categoryID,
		Images:      images,
		Active:      active,
	}
}
