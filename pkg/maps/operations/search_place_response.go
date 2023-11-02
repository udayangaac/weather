package operations

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/udayangaac/weather/pkg/maps/models"
)

func NewSearchPlaceResponse() *SearchPlaceResponse {
	return &SearchPlaceResponse{
		Body: make([]models.Place, 0),
	}
}

type SearchPlaceResponse struct {
	Body []models.Place
}

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
