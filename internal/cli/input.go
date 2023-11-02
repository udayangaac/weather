package cli

import (
	"flag"
	"fmt"
	"os"
)

// ParseInput parses the command-line input and stores the city in the provided string pointer.
func ParseInput(city *string) {
	flag.Parse()

	// Check if there are no command-line arguments, and display the help message if so.
	if len(flag.Args()) == 0 {
		help()
		os.Exit(0)
	}

	// Store the city provided as a command-line argument in the 'city' variable.
	*city = flag.Args()[0]
}

// help displays the help message for the CLI application.
func help() {
	helpMessage := `A CLI application which provides a brief weather forecast for the given city in Sweden.
Usage: ./weather [a city in Sweden]`
	fmt.Println(helpMessage)
}
