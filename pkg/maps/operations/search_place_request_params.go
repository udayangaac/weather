package operations

import (
	"net/http"
)

// SearchPlaceRequestParams represents the parameters for a place search request.
type SearchPlaceRequestParams struct {
	Country string
	City    string
}

// Build extract parameter from the HTTP request.
func (r SearchPlaceRequestParams) Build(req *http.Request) error {
	q := req.URL.Query()
	q.Add("city", r.City)
	q.Add("country", r.Country)
	req.URL.RawQuery = q.Encode()
	return nil
}
