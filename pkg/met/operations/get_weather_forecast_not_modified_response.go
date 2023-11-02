package operations

type GetWeatherForecastNotModifiedResponse struct{}

func (g GetWeatherForecastNotModifiedResponse) Error() string {
	return "weather forecast not modified"
}
