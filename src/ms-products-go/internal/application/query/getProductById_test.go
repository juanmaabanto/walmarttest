package query

import (
	"context"
	"testing"

	"github.com/juanmaabanto/go-seedwork/seedwork/errors"
	"github.com/juanmaabanto/ms-products/internal/application/response"
	"github.com/juanmaabanto/ms-products/internal/domain/products"
	"github.com/juanmaabanto/ms-products/internal/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Handle_GetProductById_Palindrome(t *testing.T) {
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

func Test_Handle_GetProductById_Not_Palindrome(t *testing.T) {
	// Arrange
	mockRepo := new(tools.MockRepository)
	ctx := context.Background()
	expected := response.ProductResponse{Id: 100, Brand: "marca", Description: "description", Image: "http://image.com", Price: 10000, ConDescuento: false}

	mockRepo.On("FindById", ctx, int64(100), mock.AnythingOfType("*products.Product")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*products.Product)
		arg.Id = 100
		arg.Brand = "marca"
		arg.Description = "description"
		arg.Image = "http://image.com"
		arg.Price = 10000
	})

	// Act
	testQuery := NewGetProductByIdHandler(mockRepo)
	result, _ := testQuery.Handle(ctx, GetProductById{Id: 100})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Equal(t, expected.Id, result.Id)
	assert.Equal(t, expected.Brand, result.Brand)
	assert.Equal(t, expected.Description, result.Description)
	assert.Equal(t, expected.Image, result.Image)
	assert.Equal(t, expected.Price, result.Price)
	assert.Equal(t, expected.ConDescuento, result.ConDescuento)
}

func Test_Handle_GetProductById_Not_Found(t *testing.T) {
	// Arrange
	mockRepo := new(tools.MockRepository)
	ctx := context.Background()

	mockRepo.On("FindById", ctx, int64(1), mock.AnythingOfType("*products.Product")).Return(errors.NewNotFoundError("No se encontro el producto por el Id")).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*products.Product)
		arg.Id = 0
	})

	// Act
	testQuery := NewGetProductByIdHandler(mockRepo)
	_, err := testQuery.Handle(ctx, GetProductById{Id: 1})

	// Asserts
	mockRepo.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "No se encontro el producto por el Id")
}
