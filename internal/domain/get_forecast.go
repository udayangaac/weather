package domain

import (
	"log"
	"time"

	"github.com/udayangaac/weather/internal/models/forecast"
)

// GetForecast fetches weather forecast for a specific city in a given country.
// It uses the provided forecastService and geoCodingService to retrieve the forecast summary.
// Parameters:
//
//	country: The name of the country.
//	city: The name of the city for which to get the weather forecast.
//	forecastService: An implementation of the ForecastService interface.
//	geoCodingService: An implementation of the GeoCodingService interface.
//
// Returns:
//
//	summary: The weather forecast summary for the given city and country.
//	nextUpdate: The duration until the next forecast update is expected.
//	err: An error, if any occurred during the retrieval process.
func GetForecast(country string, city string, forecastService ForecastService, geoCodingService GeoCodingService) (summary forecast.Summary, nextUpdate time.Duration, err error) {
	var (
		lat, lon float64
	)

	// Get the coordinates (latitude and longitude) for the provided city and country.
	lat, lon, err = geoCodingService.GetCoordByCityName(country, city)
	if err != nil {
		log.Println(err)
		return
	}

	// Retrieve the weather forecast summary for the given coordinates.
	summary, err = forecastService.GetSummary(lat, lon)

	// Set the next forecast update duration (e.g., 5 seconds).
	nextUpdate = time.Second * 5
	return
}
