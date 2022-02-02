package application

import "github.com/juanmaabanto/ms-products/internal/application/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	CreateProduct command.CreateProductHandler
}
