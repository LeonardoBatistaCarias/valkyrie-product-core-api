package validator

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/utils/constants"
)

type ProductValidator struct {
}

func (p *ProductValidator) Execute(product product.Product) error {
	checkNameConstraints(product.Name)
	checkDescriptionConstraints(product.Description)
	checkPriceConstraint(product.Price)
	checkQuantityConstraint(product.Quantity)

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

func checkPriceConstraint(price float64) bool {
	if price < constants.MIN_PRICE {

	}
	return true
}

func checkQuantityConstraint(quantity int) bool {
	if quantity < constants.MIN_QUANTITY {

	}

	return true
}
