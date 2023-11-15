package operations

import (
	"context"
	"fmt"
	"net/http"
)

const (
	geoCodeURL = "https://geocode.maps.co/search/"
)

// Client represents the interface for the geocoding client.
type Client interface {
	SearchPlace(ctx context.Context, params SearchPlaceRequestParams) (*SearchPlaceResponse, error)
}

// client is the implementation of the geocoding client.
type client struct {
	HTTPClient *http.Client
}

// NewClient creates a new geocoding client with the provided HTTP client.
func NewClient(httpClient *http.Client) Client {
	return &client{HTTPClient: httpClient}
}

// SearchPlace sends a geocoding request and returns the response.
func (c *client) SearchPlace(ctx context.Context, params SearchPlaceRequestParams) (*SearchPlaceResponse, error) {

	httpReq, err := http.NewRequest("GET", geoCodeURL, nil)
	if err != nil {
		return nil, err
	}

	if err = params.Build(httpReq); err != nil {
		return nil, err
	}

	httpResp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search places error. http status: %v", httpResp.Status)
	}

	// Create a new response object and read the HTTP response into it.
	resp := NewSearchPlaceResponse()
	if err = resp.Read(httpResp); err != nil {
		return nil, err
	}

	return resp, nil
}
