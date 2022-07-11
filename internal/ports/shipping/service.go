package shipping

import "clean/internal/entities"

type Service interface {
	GetShippingRatesForDestination(details entities.ShippingDetails) entities.ShippingRates
}
