package product

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/utils/constants"
)

type CreateProductValidator struct {
}

func (c *CreateProductValidator) Execute(p Product) error {
	checkNameConstraints(p.Name)
	checkDescriptionConstraints(p.Description)
	checkPriceConstraint(p.Price)
	checkQuantityConstraint(p.Quantity)

	return nil
}

func checkNameConstraints(name string) bool {
	length := len(name)
	if length < constants.MIN_NAME_LENGTH || length > constants.MAX_NAME_LENGTH {

	}

	return true
}

func checkDescriptionConstraints(description string) bool {
	length := len(description)
	if length < constants.MIN_DESCRIPTION_LENGTH || length > constants.MAX_DESCRIPTION_LENGTH {

	}
	return true
}

func checkPriceConstraint(price float32) bool {
	if price < constants.MIN_PRICE {

	}
	return true
}

func checkQuantityConstraint(quantity int32) bool {
	if quantity < constants.MIN_QUANTITY {

	}

	return true
}
