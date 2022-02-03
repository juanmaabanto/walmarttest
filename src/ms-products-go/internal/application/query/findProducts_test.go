package query

import (
	"context"
	"reflect"
	"testing"

	"github.com/juanmaabanto/go-seedwork/seedwork/database"
	"github.com/juanmaabanto/ms-products/internal/application/response"
	"github.com/juanmaabanto/ms-products/internal/domain/products"
	"github.com/juanmaabanto/ms-products/internal/infrastructure"
)

func TestFindProductsHandler_Handle(t *testing.T) {
	mongo_url := "mongodb+srv://root:A123a@develop.oh3sr.mongodb.net/test?retryWrites=true&w=majority"

	expected := []response.ProductResponse{
		{
			Id:           58,
			Brand:        "daad",
			Description:  "vangde oswss",
			Image:        "www.lider.cl/catalogo/images/furnitureIcon.svg",
			Price:        399362,
			ConDescuento: true,
		},
	}

	type args struct {
		query FindProducts
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		want1   []response.ProductResponse
		wantErr bool
	}{
		{"Devuelve un registro", args{FindProducts{Search: "daad", Start: 0, PageSize: 1}}, 27, expected, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := FindProductsHandler{
				repo: infrastructure.NewProductRepository(database.NewMongoConnection(context.Background(), "test", mongo_url), products.Product{}),
			}
			got, got1, err := h.Handle(context.Background(), tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindProductsHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindProductsHandler.Handle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FindProductsHandler.Handle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
