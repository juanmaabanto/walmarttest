package query

import (
	"context"
	"strconv"

	"github.com/juanmaabanto/go-seedwork/seedwork/errors"
	"github.com/juanmaabanto/ms-products/internal/application/response"
	"github.com/juanmaabanto/ms-products/internal/domain/products"
	"github.com/juanmaabanto/ms-products/internal/tools"
)

type GetProductById struct {
	Id int64
}

type GetProductByIdHandler struct {
	repo products.Repository
}

func NewGetProductByIdHandler(repo products.Repository) GetProductByIdHandler {
	if repo == nil {
		panic("nil repo")
	}

	return GetProductByIdHandler{repo: repo}
}

func (h GetProductByIdHandler) Handle(ctx context.Context, query GetProductById) (*response.ProductResponse, error) {
	result := &products.Product{}
	err := h.repo.FindById(ctx, query.Id, result)

	if result.Id == 0 {
		return nil, errors.NewNotFoundError("No se encontro el producto por el Id")
	}

	price := int64(0)
	esPalindrome := tools.EsPalindrome(strconv.Itoa(int(query.Id)))

	if esPalindrome {
		price = result.Price / 2
	} else {
		price = result.Price
	}

	response := &response.ProductResponse{
		Brand:        result.Brand,
		Description:  result.Description,
		Id:           result.Id,
		Image:        result.Image,
		Price:        price,
		ConDescuento: esPalindrome,
	}

	return response, err
}
