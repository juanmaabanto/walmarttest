package query

import (
	"context"
	"testing"

	"github.com/juanmaabanto/ms-products/internal/application/response"
	"github.com/juanmaabanto/ms-products/internal/domain/products"
	"github.com/juanmaabanto/ms-products/internal/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProductById(t *testing.T) {
	// Arrange
	mockRepo := new(tools.MockRepository)
	ctx := context.Background()
	expected := response.ProductResponse{Id: 1, Brand: "marca", Description: "description", Image: "http://image.com", Price: 5000, ConDescuento: true}

	mockRepo.On("FindById", ctx, int64(1), mock.AnythingOfType("*products.Product")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*products.Product)
		arg.Id = 1
		arg.Brand = "marca"
		arg.Description = "description"
		arg.Image = "http://image.com"
		arg.Price = 10000
	})

	// Act
	testQuery := NewGetProductByIdHandler(mockRepo)
	result, _ := testQuery.Handle(ctx, GetProductById{Id: 1})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Equal(t, expected.Id, result.Id)
	assert.Equal(t, expected.Brand, result.Brand)
	assert.Equal(t, expected.Description, result.Description)
	assert.Equal(t, expected.Image, result.Image)
	assert.Equal(t, expected.Price, result.Price)
	assert.Equal(t, expected.ConDescuento, result.ConDescuento)
}
