package operations

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/udayangaac/weather/pkg/met/models"
)

func NewGetWetherForecastResponse() *GetWetherForecastResponse {
	return &GetWetherForecastResponse{
		Body: &models.WetherForecastInfo{},
	}
}

type GetWetherForecastResponse struct {
	Body *models.WetherForecastInfo
}

func (r *GetWetherForecastResponse) Read(resp *http.Response) error {
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, r.Body)
	if err != nil {
		return err
	}
	return nil
}
