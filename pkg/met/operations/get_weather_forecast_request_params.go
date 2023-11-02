package operations

import (
	"fmt"
	"net/http"
)

// GetWeatherForecastRequestParams represents the parameters required for a weather forecast request.
type GetWeatherForecastRequestParams struct {
	Latitude, Longitude float64 // Latitude and Longitude for the location
}

// Build is a method that populates the HTTP request with the necessary query parameters and headers.
func (r GetWeatherForecastRequestParams) Build(req *http.Request) error {

	q := req.URL.Query()
	q.Add("lat", fmt.Sprintf("%.6f", r.Latitude))
	q.Add("lon", fmt.Sprintf("%.6f", r.Longitude))
	req.URL.RawQuery = q.Encode()

	// TODO: Make this configurable.
	req.Header.Add("User-Agent", "walker_udayangaac")

	return nil
}
