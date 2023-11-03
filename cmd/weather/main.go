package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/udayangaac/weather/internal/adapters"
	"github.com/udayangaac/weather/internal/cli"
	"github.com/udayangaac/weather/internal/domain"
	"github.com/udayangaac/weather/pkg/maps"
	mapsOperations "github.com/udayangaac/weather/pkg/maps/operations"
	"github.com/udayangaac/weather/pkg/met"
	metOperation "github.com/udayangaac/weather/pkg/met/operations"
	"github.com/udayangaac/weather/pkg/refreshers"
)

// country is a constant representing the default country for weather forecasts.
const country = "Sweden"

func main() {

	var city string

	// Parse the user input and store it in the 'city' variable.
	cli.ParseInput(&city)

	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, syscall.SIGTERM, syscall.SIGINT)

	forecastService := adapters.NewMetForecastService(metOperation.NewClient(met.DefaultClient(10)))
	geoCodingService := adapters.NewMapsGeoCodingService(mapsOperations.NewClient(maps.DefaultClient(10)))

	poller := refreshers.NewPoller(getCallback(city, forecastService, geoCodingService))

	poller.Poll()

	<-osSig

	poller.Stop()
}

func getCallback(city string, forecastService domain.ForecastService, geoCodingService domain.GeoCodingService) func() (time.Duration, error) {
	return func() (time.Duration, error) {
		// Get the weather forecast for the specified city and country.
		summary, duration, err := domain.GetForecast(country, city, forecastService, geoCodingService, time.Now().UTC())

		if err != nil && (err != domain.ErrNotModified || err != domain.ErrServiceBusy) {
			return duration, err
		}

		if err == domain.ErrNotModified {
			fmt.Printf("Summary not modified. Retrying in: %s.\n", duration.String())
			return duration, nil
		}

		if err == domain.ErrServiceBusy {
			fmt.Printf("Outbound services are busy. Retrying in: %s.\n", duration.String())
			return duration, nil
		}

		cli.Write(*summary)
		fmt.Printf("Weather forecast will be updated in %s.\n", duration.String())
		return duration, nil
	}
}
