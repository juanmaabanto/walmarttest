package ports

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/juanmaabanto/go-seedwork/seedwork/errors"
	"github.com/juanmaabanto/go-seedwork/seedwork/responses"
	"github.com/juanmaabanto/ms-products/internal/application"
	"github.com/juanmaabanto/ms-products/internal/application/command"
	"github.com/juanmaabanto/ms-products/internal/application/query"
	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	app application.Application
}

func NewHttpServer(application application.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

// CreateProduct godoc
// @Summary Create a new product.
// @Tags Products
// @Accept json
// @Produce json
// @Param command body products.Product true "Object to be created."
// @Success 201 {string} string "Id of the created object"
// @Failure 400 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/v1/products [post]
func (h HttpServer) AddProduct(c echo.Context) error {
	item := command.CreateProduct{}

	if err := c.Bind(&item); err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := c.Validate(item); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		panic(errors.NewValidationError(Simple(validationErrors)))
	}

	id, err := h.app.Commands.CreateProduct.Handle(c.Request().Context(), item)

	if err != nil {
		panic(err)
	}

	c.Response().Header().Set("location", c.Request().URL.String()+"/"+id)

	return c.JSON(http.StatusCreated, id)
}

// GetProduct godoc
// @Summary Get a product by Id.
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string  true  "Product Id"
// @Success 200 {object} response.ProductResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/v1/products/{id} [get]
func (h HttpServer) GetProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(errors.NewBadRequestError("El id no es válido"))
	}

	item, err := h.app.Queries.GetProductById.Handle(c.Request().Context(), query.GetProductById{Id: int64(id)})

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, item)
}

// ListProduct godoc
// @Summary Return a Product List.
// @Tags Products
// @Accept json
// @Produce json
// @Param search query string  true  "Palabra a buscar"
// @Param pageSize query int  true  "Número de resultados por página"
// @Param start query string  true  "Número de página"
// @Success 200 {object} responses.PaginatedResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/v1/products [get]
func (h HttpServer) ListProduct(c echo.Context) error {
	searchParam := c.QueryParam("search")

	if len(searchParam) < 3 {
		panic(errors.NewBadRequestError("Ingrese al menos 3 caracteres para parametro 'search'"))
	}

	pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))

	if err != nil {
		pageSize = 50
	}

	start, err := strconv.Atoi(c.QueryParam("start"))

	if err != nil {
		start = 0
	}

	total, items, err := h.app.Queries.FindProducts.Handle(c.Request().Context(), query.FindProducts{
		Search:   searchParam,
		Start:    int64(start),
		PageSize: int64(pageSize),
	})

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, responses.PaginatedResponse{
		Start:    int64(start),
		PageSize: int64(pageSize),
		Total:    total,
		Data:     items,
	})
}

func Simple(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}
