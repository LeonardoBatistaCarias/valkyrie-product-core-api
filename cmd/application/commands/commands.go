package commands

import "github.com/LeonardoBatistaCarias/valkyrie-product-core-api/cmd/application/commands/create"

type Commands struct {
	CreateProduct create.CreateProductCommandHandler
}

func NewCommands(createProduct create.CreateProductCommandHandler) *Commands {
	return &Commands{CreateProduct: createProduct}
}
