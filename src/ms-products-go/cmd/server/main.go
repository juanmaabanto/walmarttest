package main

import (
	"context"
	"log"

	_ "github.com/juanmaabanto/ms-products/docs"

	"github.com/joho/godotenv"
	"github.com/juanmaabanto/go-seedwork/seedwork/managers"
	"github.com/juanmaabanto/go-seedwork/seedwork/middleware"
	"github.com/juanmaabanto/ms-products/internal/ports"
	"github.com/juanmaabanto/ms-products/internal/service"
	"github.com/juanmaabanto/ms-products/internal/validations"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Product API
// @version v1
// @description Specifying services for micro service Product.

// @contact.name Juan Manuel Abanto Mera
// @contact.url https://www.linkedin.com/in/juanmanuelabanto/
// @contact.email jmanuelabanto@gmail.com

// @license.name MIT License
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := echo.New()
	ctx := context.Background()

	application := service.NewApplication(ctx)

	Handler(ports.NewHttpServer(application), router)
	router.Logger.Fatal(router.Start(":5500"))
}

type ServerInterface interface {
	AddProduct(c echo.Context) error
	GetProduct(c echo.Context) error
	ListProduct(c echo.Context) error
}

func Handler(si ServerInterface, router *echo.Echo) {
	if router == nil {
		router = echo.New()
	}

	router.Validator = validations.NewValidationUtil()

	loggerManager := managers.NewLoggerManager("https://mylogger.com", "ms-products")

	api := router.Group("/api/v1")

	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		LoggerErrorFunc: loggerManager.Error,
	}))

	//Swagger
	router.GET("/*", echoSwagger.WrapHandler)

	//products
	api.GET("/products", si.ListProduct)
	api.GET("/products/:id", si.GetProduct)
	api.POST("/products", si.AddProduct)
}
