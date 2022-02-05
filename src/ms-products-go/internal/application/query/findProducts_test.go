package query

import (
	"context"
	"errors"
	"testing"

	"github.com/juanmaabanto/ms-products/internal/domain/products"
	"github.com/juanmaabanto/ms-products/internal/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_NewFindProductsHandler(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic")
		}
	}()

	// The following is the code under test
	NewFindProductsHandler(nil)
}

func Test_Handle_FindProducts_Error_Count(t *testing.T) {
	// Arrange
	mockRepo := new(tools.MockRepository)
	ctx := context.Background()

	mockRepo.On("Count", ctx, mock.AnythingOfType("primitive.D")).Return(int64(0), errors.New("empty name"))

	// Act
	testQuery := NewFindProductsHandler(mockRepo)
	_, _, err := testQuery.Handle(ctx, FindProducts{Search: "t", Start: 0, PageSize: 50})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.NotNil(t, err)
}

func Test_Handle_FindProducts_Error_Paginated(t *testing.T) {
	// Arrange
	mockRepo := new(tools.MockRepository)
	ctx := context.Background()

	mockRepo.On("Count", ctx, mock.AnythingOfType("primitive.D")).Return(int64(1), nil)
	mockRepo.On("Paginated", ctx, mock.AnythingOfType("primitive.D"), mock.AnythingOfType("primitive.D"), int64(50), int64(0), mock.AnythingOfType("*[]products.Product")).Return(errors.New("empty name"))

	// Act
	testQuery := NewFindProductsHandler(mockRepo)
	_, _, err := testQuery.Handle(ctx, FindProducts{Search: "t", Start: 0, PageSize: 50})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.NotNil(t, err)
}

func Test_Handle_FindProducts_With_Palindrome(t *testing.T) {
	// Arrange
	mockRepo := new(tools.MockRepository)
	ctx := context.Background()

	mockRepo.On("Count", ctx, mock.AnythingOfType("primitive.D")).Return(int64(1), nil)
	mockRepo.On("Paginated", ctx, mock.AnythingOfType("primitive.D"), mock.AnythingOfType("primitive.D"), int64(50), int64(0), mock.AnythingOfType("*[]products.Product")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(5).(*[]products.Product)

		*arg = append(*arg, products.Product{
			Brand:       "marca",
			Description: "description",
			Image:       "image",
			Price:       10000,
		})

	})

	// Act
	testQuery := NewFindProductsHandler(mockRepo)
	total, results, _ := testQuery.Handle(ctx, FindProducts{Search: "daad", Start: 0, PageSize: 50})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Equal(t, int64(1), total)
	assert.Equal(t, int64(5000), results[0].Price)
	assert.Equal(t, true, results[0].ConDescuento)
}

func Test_Handle_FindProducts_With_Not_Palindrome(t *testing.T) {
	// Arrange
	mockRepo := new(tools.MockRepository)
	ctx := context.Background()

	mockRepo.On("Count", ctx, mock.AnythingOfType("primitive.D")).Return(int64(1), nil)
	mockRepo.On("Paginated", ctx, mock.AnythingOfType("primitive.D"), mock.AnythingOfType("primitive.D"), int64(50), int64(0), mock.AnythingOfType("*[]products.Product")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(5).(*[]products.Product)

		*arg = append(*arg, products.Product{
			Brand:       "marca",
			Description: "description",
			Image:       "image",
			Price:       10000,
		})

	})

	// Act
	testQuery := NewFindProductsHandler(mockRepo)
	total, results, _ := testQuery.Handle(ctx, FindProducts{Search: "daa", Start: 0, PageSize: 50})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Equal(t, int64(1), total)
	assert.Equal(t, int64(10000), results[0].Price)
	assert.Equal(t, false, results[0].ConDescuento)
}
