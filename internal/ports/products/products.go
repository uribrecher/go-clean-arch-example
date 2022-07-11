package products

import (
	"clean/internal/entities"
)

type Service interface {
	GetProductByID(id string) entities.Product
	GetProductByBarcode(barcode string) entities.Product
	SaveProduct(product entities.Product)
}
