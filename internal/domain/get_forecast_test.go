package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/udayangaac/weather/internal/models/forecast"
)

func TestGetForecast(t *testing.T) {
	// Create an instance of the MockGeoCodingService and MockForecastService.
	expectedSummary := &forecast.Summary{
		Title:   "Weather forecast summary.",
		Country: "Sri Lanka",
		City:    "Colombo",
		Expires: "Mon, 02 Jan 2024 11:20:00 GMT",
	}
	expectedNextUpdateSecond := float64(3600)

	geoCodingService := &MockGeoCodingService{
		Latitude:  5.00,
		Longitude: 81.9092,
		Error:     nil,
	}

	forecastService := &MockForecastService{
		Summary: expectedSummary,
		Error:   nil,
	}

	currentTime, _ := time.Parse(time.RFC1123, "Mon, 02 Jan 2024 10:20:00 GMT")
	actualSummary, actualNextUpdate, actualErr := GetForecast("Sri Lanka", "Colombo", forecastService, geoCodingService, currentTime)

	assert.Equal(t, expectedSummary, actualSummary, "Test the summary")
	assert.Equal(t, expectedNextUpdateSecond, actualNextUpdate.Seconds(), "Test the next update in seconds")
	assert.Nil(t, actualErr, "Test actual error")
}

// MockGeoCodingService is a mock implementation of the GeoCodingService interface.
type MockGeoCodingService struct {
	Latitude, Longitude float64
	Error               error
}

func (m *MockGeoCodingService) GetCoordByCityName(country string, city string) (float64, float64, error) {
	return m.Latitude, m.Longitude, m.Error
}

// MockForecastService is a mock implementation of the ForecastService interface.
type MockForecastService struct {
	Summary *forecast.Summary
	Error   error
}

func (m *MockForecastService) GetSummary(lat float64, lon float64) (*forecast.Summary, error) {
	return m.Summary, m.Error
}
