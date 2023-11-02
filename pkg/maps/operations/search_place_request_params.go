package operations

import (
	"net/http"
)

type SearchPlaceRequestParams struct {
	Country, City string
}

func (r SearchPlaceRequestParams) Build(req *http.Request) error {
	q := req.URL.Query()
	q.Add("city", r.City)
	q.Add("country", r.Country)
	req.URL.RawQuery = q.Encode()
	return nil
}
