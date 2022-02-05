# walmarttest
Backend desarrollado con Go 1.17, y frontend desarrollado con React JS, MongoDB, Clean Architecture con CQRS

- microservice: src/ms-products-go
- Frontend: src/lider-app

Uso de librería propia para abstraer implementaciones comunes, repositorios genericos, middleware para el logging, etc. dejo enlace del repositorio
https://github.com/juanmaabanto/go-seedwork

## Instalación

### Requisitos

  1. Instalar docker engine
  2. Instalar docker-compose
  3. Tener disponibles puertos 3000 y 5500

### Pasos

  1. Ejecutar docker compose
  
     docker-compose up -d
     
  2. Backend: ir a http://localhost:5500 para ver la especificación swagger
  
  3. Frontend: ir a http://localhost:3000



## Pruebas Unitarias

El backend ha sido desarrollado con Go

### Requisitos

  1. Instalar Go https://go.dev/dl/

### Ejecución

  ```sh
// ir a proyecto
cd src/ms-products-go

// descargar dependencias
go mod tidy

// ejecutar pruebas
go test ./internal/application/... -v

// cobertura de pruebas
go test ./internal/application/... -cover
```
