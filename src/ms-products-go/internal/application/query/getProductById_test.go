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

var (
	mongo_url = "mongodb+srv://root:A123a@develop.oh3sr.mongodb.net/test?retryWrites=true&w=majority"
)

func TestGetProductByIdHandler_Handle(t *testing.T) {
	expected := &response.ProductResponse{
		Id:           100,
		Brand:        "qeiydij",
		Description:  "cxzbz lahbhe",
		Image:        "www.lider.cl/catalogo/images/computerIcon.svg",
		Price:        756530,
		ConDescuento: false,
	}

	type args struct {
		query GetProductById
	}
	tests := []struct {
		name    string
		args    args
		want    *response.ProductResponse
		wantErr bool
	}{
		{"Buscar Id", args{query: GetProductById{100}}, expected, false},
		{"Error cuando no encuentra Id", args{query: GetProductById{132300}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := GetProductByIdHandler{
				repo: infrastructure.NewProductRepository(database.NewMongoConnection(context.Background(), "test", mongo_url), products.Product{}),
			}
			got, err := h.Handle(context.Background(), tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductByIdHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByIdHandler.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
