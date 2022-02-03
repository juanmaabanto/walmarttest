FROM golang:1.17-alpine

WORKDIR /app

COPY src/ms-products-go/go.mod ./
COPY src/ms-products-go/go.sum ./
RUN go mod download

# Copy the code into the container
COPY src/ms-products-go .

# Build the application
RUN go build -o main ./cmd/server

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /app/main .
RUN cp /app/.env .

# Export necessary port
EXPOSE 5500

# Command to run when starting the container
CMD ["/dist/main"]