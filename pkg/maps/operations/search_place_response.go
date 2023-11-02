package operations

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/udayangaac/weather/pkg/maps/models"
)

// NewSearchPlaceResponse creates a new SearchPlaceResponse with an empty Body slice.
func NewSearchPlaceResponse() *SearchPlaceResponse {
	return &SearchPlaceResponse{
		Body: make([]models.Place, 0),
	}
}

// SearchPlaceResponse represents the response structure for searching places.
type SearchPlaceResponse struct {
	Body []models.Place
}

// Read reads and processes the response from an HTTP response.
func (r *SearchPlaceResponse) Read(resp *http.Response) error {
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &r.Body)
	if err != nil {
		return err
	}
	return nil
}
