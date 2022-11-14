package validator

import "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/domain/product"

type Validator interface {
	Execute(p product.Product) error
}
