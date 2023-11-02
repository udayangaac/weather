package operations

type GetWeatherForecastTooManyRequestsResponse struct{}

func (g GetWeatherForecastTooManyRequestsResponse) Error() string {
	return "too many requests"
}
