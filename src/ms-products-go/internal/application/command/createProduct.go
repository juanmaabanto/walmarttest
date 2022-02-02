package command

import (
	"context"
	"time"

	"github.com/juanmaabanto/go-seedwork/seedwork/errors"
	"github.com/juanmaabanto/ms-products/internal/domain/products"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateProduct struct {
	Brand       string `json:"brand" validate:"required,max=20"`
	Description string `json:"description" validate:"required,max=50"`
	Image       string `json:"image" validate:"required,uri"`
	Price       int64  `json:"price" validate:"required,number,gt=0"`
}

type CreateProductHandler struct {
	repo products.Repository
}

func NewCreateProductHandler(repo products.Repository) CreateProductHandler {
	if repo == nil {
		panic("nil repo")
	}

	return CreateProductHandler{repo: repo}
}

func (h CreateProductHandler) Handle(ctx context.Context, command CreateProduct) (string, error) {
	count, err := h.repo.Count(ctx, bson.M{"description": primitive.Regex{
		Pattern: "^" + command.Description + "$",
		Options: "i",
	}})

	if err != nil {
		return "", err
	}

	if count > 0 {
		return "", errors.NewBadRequestError("Ya existe un elemento con el mismo nombre")
	}

	item := products.Product{}

	item.Brand = command.Brand
	item.Description = command.Description
	item.Image = command.Image
	item.Price = command.Price
	item.CreatedAt = time.Now()
	item.CreatedBy = "admin"

	id, err := h.repo.InsertOne(ctx, item)

	if err != nil {
		return id, err
	}

	return id, nil
}
