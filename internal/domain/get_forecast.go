package domain

import (
	"time"

	"github.com/udayangaac/weather/internal/models/forecast"
)

// GetForecast fetches the weather forecast for a specific city in a given country.
// It utilizes the provided forecastService and geoCodingService to obtain the forecast summary.
//
// Parameters:
//
//	country: The name of the country.
//	city: The name of the city for which to get the weather forecast.
//	forecastService: An implementation of the ForecastService interface.
//	geoCodingService: An implementation of the GeoCodingService interface.
//	currentTime: The current time used to calculate the next forecast update duration.
//
// Returns:
//
//	summary: The weather forecast summary for the given city and country.
//	nextUpdate: The duration until the next forecast update is expected.
//	err: An error, if any occurred during the retrieval process.
func GetForecast(country string, city string, forecastService ForecastService, geoCodingService GeoCodingService, currentTime time.Time) (summary *forecast.Summary, nextUpdate time.Duration, err error) {
	var (
		lat, lon float64
	)

	// Set the next forecast update duration (e.g., 5 seconds).
	nextUpdate = 5 * time.Second

	// Get the coordinates (latitude and longitude) for the provided city and country.
	lat, lon, err = geoCodingService.GetCoordByCityName(country, city)
	if err != nil {
		return
	}

	// Retrieve the weather forecast summary for the given coordinates.
	summary, err = forecastService.GetSummary(lat, lon)
	if err != nil {
		return
	}

	expiryTime, err := time.Parse(time.RFC1123, summary.Expires)
	if err != nil {
		return
	}

	summary.Title = "Weather forecast summary."
	summary.City = city
	summary.Country = country

	nextUpdate = expiryTime.Sub(currentTime)
	return
}
