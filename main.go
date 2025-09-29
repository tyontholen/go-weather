package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"weather/weather"
)

var cityCoords = map[string][2]float64{
	"london":    {51.5072, -0.1276},
	"stockholm": {59.3293, 18.0686},
	"newyork":   {40.7128, -74.0060},
	"tokyo":     {35.6762, 139.6503},
	"oslo":      {59.9139, 10.7522},
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println(" weather <city>			(e.g weather Stockholm)")
	fmt.Println(" weather <latitude> <lon>	(e.g weather 59.3293 18.0686)")
}

func parseInputs(args []string) (float64, float64, error) {
	// if single argument, try city lookup

	if len(args) == 1 {
		city := strings.ToLower(args[0])
		if coords, ok := cityCoords[city]; ok {
			return coords[0], coords[1], nil
		}
		return 0, 0, fmt.Errorf("Error: Unknown city: %s", args[0])
	}

	// if two or more args, try parse the first 2 as latitude & longitude
	if len(args) >= 2 {
		lat, err1 := strconv.ParseFloat(args[0], 64)
		lon, err2 := strconv.ParseFloat(args[1], 64)
		if err1 == nil && err2 == nil {
			return lat, lon, nil
		}

		// if parsing fails, maybe user types multi-word city, like New York
		city := strings.ToLower(strings.Join(args, ""))
		if coords, ok := cityCoords[city]; ok {
			return coords[0], coords[1], nil
		}
		return 0, 0, fmt.Errorf("Error: Could njot parse coordinates or find city")

	}
	return 0, 0, fmt.Errorf("Error: Not enough arguments")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	lat, lon, err := parseInputs(os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		usage()
		return
	}

	// create the client - returned pointer

	client := weather.NewClient()

	// call method with pointer reciever

	cw, err := client.GetCurrentWeather(lat, lon)
	if err != nil {
		fmt.Println("Error fetching weather", err)
		return
	}

	fmt.Printf("Location: %.4f, %.4f\n", lat, lon)
	fmt.Printf("Time: %s\n", cw.Time)
	fmt.Printf("Temperature: %.1f °C\n", cw.Temperature)
	fmt.Printf("Wind speed: %.1f m/s\n", cw.WindSpeed)
	fmt.Printf("Weather code: %d\n", cw.WeatherCode)

}
