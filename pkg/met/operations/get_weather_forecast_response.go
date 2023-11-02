package operations

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/udayangaac/weather/pkg/met/models"
)

// NewGetWeatherForecastResponse initializes a new GetWeatherForecastResponse object
func NewGetWeatherForecastResponse() *GetWeatherForecastResponse {
	return &GetWeatherForecastResponse{
		Body: &models.WeatherForecastInfo{},
	}
}

// GetWeatherForecastResponse represents the response structure for weather forecasts
type GetWeatherForecastResponse struct {
	Body *models.WeatherForecastInfo
}

// Read reads and processes the HTTP response body
func (r *GetWeatherForecastResponse) Read(resp *http.Response) error {
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, r.Body)
	if err != nil {
		return err
	}

	return nil
}
