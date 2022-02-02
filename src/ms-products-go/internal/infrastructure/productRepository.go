package infrastructure

import (
	"github.com/juanmaabanto/go-seedwork/seedwork"
	"github.com/juanmaabanto/go-seedwork/seedwork/database"
	"github.com/juanmaabanto/ms-products/internal/domain/products"
)

type ProductRepository struct {
	seedwork.BaseRepository
}

func NewProductRepository(connection database.MongoConnection, document products.Product) ProductRepository {
	repository := ProductRepository{
		BaseRepository: *seedwork.NewBaseRepository(connection, &document),
	}

	return repository
}
