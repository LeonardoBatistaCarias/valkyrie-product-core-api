package model

import uuid "github.com/satori/go.uuid"

type ImageCommand struct {
	Name      string
	ProductID uuid.UUID
}

func NewImageCommand(name string, productID uuid.UUID) *ImageCommand {
	return &ImageCommand{Name: name, ProductID: productID}
}
