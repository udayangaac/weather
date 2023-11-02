package operations

import (
	"context"
	"net/http"
)

// Define the URL for the weather forecast API.
const locationForecastCompactURL = "https://api.met.no/weatherapi/locationforecast/2.0/compact/"

// Client is an interface for interacting with the weather forecast service.
type Client interface {
	GetWeatherForecast(ctx context.Context, params GetWeatherForecastRequestParams) (*GetWeatherForecastResponse, error)
}

// client is the implementation of the Client interface.
type client struct {
	HttpClient *http.Client
}

// NewClient creates a new weather forecast client with the given HTTP client.
func NewClient(httpClient *http.Client) Client {
	return &client{HttpClient: httpClient}
}

// GetWeatherForecast retrieves weather forecast data using the provided parameters.
func (c *client) GetWeatherForecast(ctx context.Context, params GetWeatherForecastRequestParams) (*GetWeatherForecastResponse, error) {
	httpReq, err := http.NewRequest("GET", locationForecastCompactURL, nil)
	if err != nil {
		return nil, err
	}

	if err = params.Build(httpReq); err != nil {
		return nil, err
	}

	httpResp, err := c.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	resp := NewGetWeatherForecastResponse()
	if err = resp.Read(httpResp); err != nil {
		return nil, err
	}

	return resp, nil
}
