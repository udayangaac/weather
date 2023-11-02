package operations

import (
	"context"
	"net/http"
)

// Define the URL for the weather forecast API.
const locationForecastCompactURL = "https://api.met.no/weatherapi/locationforecast/2.0/compact/"

const (
	ifModifiedSinceHeaderKey = "If-Modified-Since"
	lastModifiedHeaderKey    = "Last-Modified"
	expiresHeaderKey         = "Expires"
	userAgentHeaderKey       = "User-Agent"

	maximumRetryCount = 5
)

// Client is an interface for interacting with the weather forecast service.
type Client interface {
	GetWeatherForecast(ctx context.Context, params GetWeatherForecastRequestParams) (*GetWeatherForecastResponse, error)
}

// client is the implementation of the Client interface.
type client struct {
	lastModified string
	HttpClient   *http.Client
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

	if c.lastModified != "" {
		httpReq.Header.Add(ifModifiedSinceHeaderKey, c.lastModified)
	}

	httpResp, err := c.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode == http.StatusNotModified {
		return nil, &GetWeatherForecastNotModifiedResponse{}
	}

	if httpResp.StatusCode == http.StatusTooManyRequests {
		return nil, &GetWeatherForecastNotModifiedResponse{}
	}

	if lastModifiedStr := httpResp.Header.Get(lastModifiedHeaderKey); lastModifiedStr != "" {
		c.lastModified = lastModifiedStr
	}

	resp := NewGetWeatherForecastResponse()
	if err = resp.Read(httpResp); err != nil {
		return nil, err
	}

	return resp, nil
}
