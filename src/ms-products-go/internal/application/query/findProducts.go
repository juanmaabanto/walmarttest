package query

import (
	"context"

	"github.com/juanmaabanto/ms-products/internal/application/response"
	"github.com/juanmaabanto/ms-products/internal/application/tools"
	"github.com/juanmaabanto/ms-products/internal/domain/products"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindProducts struct {
	Search   string
	Start    int64
	PageSize int64
}

type FindProductsHandler struct {
	repo products.Repository
}

func NewFindProductsHandler(repo products.Repository) FindProductsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return FindProductsHandler{repo: repo}
}

func (h FindProductsHandler) Handle(ctx context.Context, query FindProducts) (int64, []response.ProductResponse, error) {
	var items []products.Product
	results := []response.ProductResponse{}

	filter := bson.D{
		{"$or",
			bson.A{
				bson.D{{"brand", primitive.Regex{
					Pattern: query.Search,
					Options: "i",
				}}},
				bson.D{{"description", primitive.Regex{
					Pattern: query.Search,
					Options: "i",
				}}},
			}},
	}

	total, err := h.repo.Count(ctx, filter)

	if err != nil {
		return 0, results, err
	}

	err = h.repo.Paginated(ctx, filter, bson.D{}, query.PageSize, query.Start, &items)

	if err != nil {
		return 0, results, err
	}

	esPalindrome := tools.EsPalindrome(query.Search)

	for _, element := range items {
		price := int64(0)

		if esPalindrome {
			price = element.Price / 2
		} else {
			price = element.Price
		}

		results = append(results, response.ProductResponse{
			Id:           element.Id,
			Brand:        element.Brand,
			Description:  element.Description,
			Image:        element.Image,
			Price:        price,
			ConDescuento: esPalindrome,
		})
	}

	return total, results, err
}
