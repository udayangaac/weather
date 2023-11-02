package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/evertras/bubble-table/table"
	"github.com/udayangaac/weather/internal/models/forecast"
)

// Column keys for the table
const (
	dateColumnKey                string = "date"
	timeColumnKey                string = "time"
	temperatureColumnKey         string = "temperature"
	windSpeedColumnKey           string = "windSpeed"
	precipitationAmountColumnKey string = "precipitationAmount"
)

// getColumns creates table columns based on forecast header
func getColumns(header forecast.Header) []table.Column {
	return []table.Column{
		table.NewColumn(dateColumnKey, header.Date, len(header.Date)),
		table.NewColumn(timeColumnKey, header.Time, len(header.Time)),
		table.NewColumn(temperatureColumnKey, header.Temperature, len(header.Temperature)),
		table.NewColumn(windSpeedColumnKey, header.WindSpeed, len(header.Temperature)),
		table.NewColumn(precipitationAmountColumnKey, header.PrecipitationAmount, len(header.PrecipitationAmount)),
	}
}

// getRows creates table rows based on forecast rows
func getRows(rows []forecast.Row) []table.Row {
	tableRows := make([]table.Row, 0)
	for _, row := range rows {
		// Create a new row with the corresponding column data
		row := table.NewRow(table.RowData{
			dateColumnKey:                row.Date,
			timeColumnKey:                row.Time,
			temperatureColumnKey:         row.Temperature,
			windSpeedColumnKey:           row.WindSpeed,
			precipitationAmountColumnKey: row.PrecipitationAmount,
		})
		tableRows = append(tableRows, row)
	}
	return tableRows
}

// Write prints the forecast summary to the console
func Write(summary forecast.Summary) {
	t := table.New(getColumns(summary.Header)).WithRows(getRows(summary.Rows))
	body := strings.Builder{}
	body.WriteString(fmt.Sprintf("Updated @ %v\n", time.Now().Format(time.DateTime)))
	body.WriteString(t.View())
	fmt.Println(body.String())
}
