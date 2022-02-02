package products

import "github.com/juanmaabanto/go-seedwork/seedwork"

type Product struct {
	Brand             string `bson:"brand"`
	Description       string `bson:"description"`
	Image             string `bson:"image"`
	Price             int64  `bson:"price"`
	seedwork.Document `bson:"inline"`
}

func (_ Product) GetCollectionName() string {
	return "products"
}
