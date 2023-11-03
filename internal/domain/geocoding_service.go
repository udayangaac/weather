package domain

// GeoCodingService is an interface for obtaining coordinates (latitude and longitude) for a given city and country.
type GeoCodingService interface {
	GetCoordByCityName(country, city string) (latitude, longitude float64, err error)
}
