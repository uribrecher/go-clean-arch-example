package customers

import "clean/internal/entities"

type Service interface {
	GetCustomer(id string) entities.Customer
	SaveCustomer(customer entities.Customer)
}
