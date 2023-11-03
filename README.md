# weather
A CLI application which provides a brief weather forecast for the given city in Sweden.

## Usage

To use the Weather CLI application, simply run the following command in your terminal:
```bash
./weather [city in Sweden]
```
Replace `[city in Sweden]` with the name of the city for which you want to get the weather forecast.

## Installation
1. Make sure you have Go installed on your system.
2. Clone the repository:

```bash
git clone git@github.com:udayangaac/weather.git
```
3. Navigate to the project directory:

```bash
cd weather
```
4. Build the CLI application:

```bash
# Create bin directory
mkdir bin

# Create binary file
go build  -o bin/weather cmd/weather/main.go 
```

5. Run the application:

```bash
./bin/weather [city in Sweden]
```
__Example__:  
For example, to get the weather forecast for Stockholm:

```bash
./bin/weather Stockholm
```
The application will then provide a brief weather summary for Stockholm.

## Dependencies

This CLI application relies on external weather data sources and may require network access to fetch weather information.

## License

This project is licensed under the BSD 2-Clause -see the LICENSE file for details.