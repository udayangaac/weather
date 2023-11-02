package domain

import (
	"errors"

	"github.com/udayangaac/weather/internal/models/forecast"
)

var (
	ErrNotModified = errors.New("not modified")
)

// ForecastService is an interface that defines methods for retrieving weather forecasts.
type ForecastService interface {
	// GetSummary retrieves a weather forecast summary for the specified latitude and longitude.
	GetSummary(latitude, longitude float64) (forecast.Summary, error)
}
