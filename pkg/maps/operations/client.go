package operations

import (
	"context"
	"net/http"
)

const (
	geoCodeUrl = "https://geocode.maps.co/search/"
)

type Client interface {
	SearchPlace(ctx context.Context, params SearchPlaceRequestParams) (*SearchPlaceResponse, error)
}

type client struct {
	HttpClient *http.Client
}

func NewClient(httpClient *http.Client) Client {
	return &client{HttpClient: httpClient}
}

func (c *client) SearchPlace(ctx context.Context, params SearchPlaceRequestParams) (*SearchPlaceResponse, error) {
	httpReq, err := http.NewRequest("GET", geoCodeUrl, nil)
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

	resp := NewSearchPlaceResponse()
	if err = resp.Read(httpResp); err != nil {
		return nil, err
	}

	return resp, nil
}
