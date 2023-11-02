package adapters

import (
	"context"
	"fmt"
	"time"

	"github.com/udayangaac/weather/internal/domain"
	"github.com/udayangaac/weather/internal/models/forecast"
	"github.com/udayangaac/weather/pkg/met/operations"
)

// NewMetForecastService creates a new MetForecastService using the provided client.
func NewMetForecastService(client operations.Client) domain.ForecastService {
	return &forecastService{Client: client}
}

// forecastService is an implementation of the ForecastService interface for the MET service.
type forecastService struct {
	Client operations.Client
}

// GetSummary retrieves weather forecast summary based on latitude and longitude.
func (f *forecastService) GetSummary(latitude, longitude float64) (forecast.Summary, error) {
	var summary forecast.Summary
	resp, err := f.Client.GetWeatherForecast(context.Background(), operations.GetWeatherForecastRequestParams{
		Latitude:  latitude,
		Longitude: longitude,
	})

	if err != nil {
		return summary, err
	}

	summary = forecast.Summary{
		Header: forecast.Header{
			Date:                "Date (UTC)",
			Time:                "Time (UTC)",
			Temperature:         fmt.Sprintf("Temperature (%v)", resp.Body.Properties.Meta.Units.AirTemperature),
			WindSpeed:           fmt.Sprintf("Wind Speed (%v)", resp.Body.Properties.Meta.Units.WindSpeed),
			PrecipitationAmount: fmt.Sprintf("Precipitation (%v)", resp.Body.Properties.Meta.Units.PrecipitationAmount),
		},
		Rows: make([]forecast.Row, 0),
	}

	for _, val := range resp.Body.Properties.Timeseries {
		row := forecast.Row{
			Date:                val.Time.Format(time.DateOnly),
			Time:                val.Time.Format(time.TimeOnly),
			Temperature:         val.Data.Instant.Details.AirTemperature,
			WindSpeed:           val.Data.Instant.Details.WindSpeed,
			PrecipitationAmount: val.Data.Instant.Details.PrecipitationAmount,
		}
		summary.Rows = append(summary.Rows, row)
	}

	return summary, nil
}
