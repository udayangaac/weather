package operations

import (
	"fmt"
	"net/http"
)

type GetWetherForecastRequestParams struct {
	Latitude, Longitude float64
}

func (r GetWetherForecastRequestParams) Build(req *http.Request) error {
	q := req.URL.Query()
	q.Add("lat", fmt.Sprintf("%.6f", r.Latitude))
	q.Add("lon", fmt.Sprintf("%.6f", r.Longitude))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("User-Agent", "walker_udayangaac")
	return nil
}
