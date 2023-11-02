package operations

import (
	"context"
	"net/http"
)

const (
	locationForecastCompactURL = "https://api.met.no/weatherapi/locationforecast/2.0/compact/"
)

type Client interface {
	GetWhetherForecast(ctx context.Context, params GetWetherForecastRequestParams) (*GetWetherForecastResponse, error)
}

type client struct {
	HttpClient *http.Client
}

func NewClient(httpClient *http.Client) Client {
	return &client{HttpClient: httpClient}
}

func (c *client) GetWhetherForecast(ctx context.Context, params GetWetherForecastRequestParams) (*GetWetherForecastResponse, error) {
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

	resp := NewGetWetherForecastResponse()
	if err = resp.Read(httpResp); err != nil {
		return nil, err
	}

	return resp, nil
}
