package application

import (
	"github.com/juanmaabanto/ms-products/internal/application/command"
	"github.com/juanmaabanto/ms-products/internal/application/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateProduct command.CreateProductHandler
}

type Queries struct {
	GetProductById query.GetProductByIdHandler
}
