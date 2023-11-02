package forecast

// Summary represents the weather forecast summary.
type Summary struct {
	Title      string // Title of the forecast summary.
	ExpiryTime string // Expiry time of the forecast.
	Header     Header
	Rows       []Row
}

// Header represents the header information for the weather forecast.
// The fields contain units for relevant data.
type Header struct {
	Date                string
	Time                string
	Temperature         string
	WindSpeed           string
	PrecipitationAmount string
}

// Row represents a row of weather forecast data.
type Row struct {
	Date                string
	Time                string
	Temperature         float64
	WindSpeed           float64
	PrecipitationAmount float64
}
