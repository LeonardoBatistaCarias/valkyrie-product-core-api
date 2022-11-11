package product

import uuid "github.com/satori/go.uuid"

type ProductImage struct {
	Name      string
	ProductID uuid.UUID
}
