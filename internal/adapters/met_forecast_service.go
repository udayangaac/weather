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
func (f *forecastService) GetSummary(latitude, longitude float64) (*forecast.Summary, error) {
	resp, err := f.Client.GetWeatherForecast(context.Background(), operations.GetWeatherForecastRequestParams{
		Latitude:  latitude,
		Longitude: longitude,
	})

	if err != nil {
		if _, ok := err.(*operations.GetWeatherForecastNotModifiedResponse); ok {
			return nil, domain.ErrNotModified
		}
		return nil, err
	}

	summary := forecast.Summary{
		Expires: resp.Expires,
		Header: forecast.Header{
			Date:                "Date (UTC)",
			Time:                "Time (UTC)",
			Temperature:         formatUnit("Temperature", resp.Body.Properties.Meta.Units.AirTemperature),
			WindSpeed:           formatUnit("Wind Speed", resp.Body.Properties.Meta.Units.WindSpeed),
			PrecipitationAmount: formatUnit("Precipitation", resp.Body.Properties.Meta.Units.PrecipitationAmount),
		},
		Rows: make([]forecast.Row, len(resp.Body.Properties.Timeseries)),
	}

	for i, val := range resp.Body.Properties.Timeseries {
		summary.Rows[i] = forecast.Row{
			Date:                val.Time.Format(time.DateOnly),
			Time:                val.Time.Format(time.TimeOnly),
			Temperature:         val.Data.Instant.Details.AirTemperature,
			WindSpeed:           val.Data.Instant.Details.WindSpeed,
			PrecipitationAmount: val.Data.Instant.Details.PrecipitationAmount,
		}
	}

	return &summary, nil
}

func formatUnit(name, unit string) string {
	return fmt.Sprintf("%s (%v)", name, unit)
}
