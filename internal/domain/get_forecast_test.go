package domain

import (
	"testing"
	"time"

	"github.com/udayangaac/weather/internal/models/forecast"
)

// MockGeoCodingService is a mock implementation of the GeoCodingService interface.
type MockGeoCodingService struct{}

func (m *MockGeoCodingService) GetCoordByCityName(country string, city string) (float64, float64, error) {
	return 40.7128, -74.0060, nil
}

// MockForecastService is a mock implementation of the ForecastService interface.
type MockForecastService struct{}

func (m *MockForecastService) GetSummary(lat float64, lon float64) (*forecast.Summary, error) {
	return &forecast.Summary{}, nil
}

func TestGetForecast(t *testing.T) {
	// Create an instance of the MockGeoCodingService and MockForecastService.
	geoCodingService := &MockGeoCodingService{}
	forecastService := &MockForecastService{}

	// Define a current time for testing.
	currentTime := time.Now()

	// Call the GetForecast function with mock services and current time.
	summary, nextUpdate, err := GetForecast("USA", "New York", forecastService, geoCodingService, currentTime)
	if err != nil {
		t.Errorf("GetForecast failed with error: %v", err)
	}

	if summary == nil {
		t.Error("GetForecast returned a nil summary")
	}
	if nextUpdate < 0 || nextUpdate > time.Hour {
		t.Errorf("Invalid nextUpdate value: %v", nextUpdate)
	}
}
