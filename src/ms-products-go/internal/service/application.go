package service

import (
	"context"
	"os"

	"github.com/juanmaabanto/go-seedwork/seedwork/database"
	"github.com/juanmaabanto/ms-products/internal/application"
	"github.com/juanmaabanto/ms-products/internal/application/command"
	"github.com/juanmaabanto/ms-products/internal/domain/products"
	"github.com/juanmaabanto/ms-products/internal/infrastructure"
)

func NewApplication(ctx context.Context) application.Application {
	conn := database.NewMongoConnection(ctx, os.Getenv("MONGODB_NAME"), os.Getenv("MONGODB_URI"))
	document := new(products.Product)

	productRepository := infrastructure.NewProductRepository(conn, *document)

	return application.Application{
		Commands: application.Commands{
			CreateProduct: command.NewCreateProductHandler(productRepository),
		},
	}
}
