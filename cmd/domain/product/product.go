package product

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/infrastructure/product/proto/pb/model"
	time_converter "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/time"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Product struct {
	ProductID   uuid.UUID `json:"product_id" validate:"required"`
	Name        string    `json:"name" validate:"required,gte=5,lte=50"`
	Description string    `json:"description" validate:"required,gte=20,lte=400"`
	Brand       Brand     `json:"brand" validate:"required"`
	Price       float32   `json:"price" validate:"required,gte=0"`
	Quantity    int32     `json:"quantity" validate:"required,gte=0"`
	CategoryID  uuid.UUID `json:"category_id" validate:"required,gte=0"`
	//ProductImages []*Image  `json:"product_images" validate:"required,gte=0"`
	Active    bool
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func NewProduct(
	productID uuid.UUID,
	name string,
	description string,
	brand Brand,
	price float32,
	quantity int32,
	categoryID uuid.UUID,
	//images []*Image,
	active bool,
	createdAt time.Time) *Product {
	return &Product{
		ProductID:   productID,
		Name:        name,
		Description: description,
		Brand:       brand,
		Price:       price,
		Quantity:    quantity,
		CategoryID:  categoryID,
		//ProductImages: images,
		Active:    active,
		CreatedAt: createdAt,
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}

func ProductFromGrpcMessage(p model.Product) *Product {
	return &Product{
		ProductID:   uuid.FromStringOrNil(p.ProductID),
		Name:        p.Name,
		Description: p.Description,
		Brand:       Brand(p.Brand),
		Price:       p.Price,
		Quantity:    p.Quantity,
		CategoryID:  uuid.FromStringOrNil(p.CategoryID),
		//ProductImages: images,
		Active:    p.Active,
		CreatedAt: *time_converter.TimeRFC3339From(p.CreatedAt),
		UpdatedAt: time_converter.TimeRFC3339From(p.UpdatedAt),
		DeletedAt: time_converter.TimeRFC3339From(p.DeletedAt),
	}
}

func (p *Product) Deactivate() {
	actualDate := time.Now()
	if p.DeletedAt == nil || p.DeletedAt.Before(p.CreatedAt) {
		p.DeletedAt = &actualDate
	}
	p.Active = false
	p.UpdatedAt = &actualDate
}
