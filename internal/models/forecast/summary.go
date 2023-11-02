package forecast

type Summary struct {
	Title      string
	ExpiryTime string
	Header     Header
	Rows       []Row
}

type Header struct {
	Date                string
	Time                string
	Temperature         string
	WindSpeed           string
	PrecipitationAmount string
}

type Row struct {
	Date                string
	Time                string
	Temperature         float64
	WindSpeed           float64
	PrecipitationAmount float64
}
