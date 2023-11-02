package adapters

import (
	"context"
	"strconv"

	"github.com/udayangaac/weather/internal/domain"
	"github.com/udayangaac/weather/pkg/maps/operations"
)

// NewMapsGeoCodingService creates a new GeoCodingService using the provided client.
func NewMapsGeoCodingService(client operations.Client) domain.GeoCodingService {
	return &geoCodingService{Client: client}
}

// geoCodingService is an implementation of the GeoCodingService interface.
type geoCodingService struct {
	Client operations.Client
}

// GetCoordByCityName returns the latitude and longitude of a city in a specific country.
func (f *geoCodingService) GetCoordByCityName(country, city string) (latitude, longitude float64, err error) {
	resp, err := f.Client.SearchPlace(context.Background(), operations.SearchPlaceRequestParams{
		Country: country,
		City:    city,
	})

	if err != nil {
		return 0, 0, err
	}

	if len(resp.Body) == 0 {
		return 0, 0, domain.ErrLocationNotFound
	}

	latitude, err = strconv.ParseFloat(resp.Body[0].Lat, 64)
	if err != nil {
		return 0, 0, err
	}

	longitude, err = strconv.ParseFloat(resp.Body[0].Lon, 64)
	if err != nil {
		return 0, 0, err
	}

	return latitude, longitude, nil
}
