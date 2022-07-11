package entities

type ShippingDetails interface {
	GetCity() string
	GetCountryCode() string
	GetStreetName() string
	GetBuildingNumber() string
	GetApartmentID() string
	GetFirstName() string
	GetLastName() string
}

type ShippingRates interface {
	GetRateByWeight(weight string) string
	GetRateByVolume(volume string) string
}
