package orders

import "clean/internal/entities"

type Service interface {
	NewOrder() entities.Order
	SaveOrder(entities.Order)
}
