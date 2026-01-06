package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	

	"github.com/Krish-jswl/weatherfetch/internal/weatherapi"
)

func shortLoc(loc string) string {
	parts := strings.Split(loc, ",")
	city := strings.TrimSpace(parts[1])
	metro := strings.TrimSpace(parts[3])

	short := city + ", " + metro
	return short
}

func main() {

	args := os.Args[1:]
	
	if len(args) < 1 {
		println("use: weatherfetch -p 123456")
		return
	}

	var pin int

	if args[0] == "-p" {
		tpin, err := strconv.Atoi(args[1])
		pin = tpin;
		if err != nil {
			fmt.Println("Enter a valid pin")
			return
		}
	}

	strconv.Atoi(args[0])

	strpin := strconv.Itoa(pin)

	location, err := weatherapi.GeoPincode(strpin)
	if err != nil {
		print(err)
	}

	weather, err := weatherapi.GetCurrentWeather(location.Lat, location.Lon)


	locname := shortLoc(location.DisplayName)

	
	fmt.Printf(`
      ☁️  WeatherFetch
      ───────────────
      City        : %s
      Temp        : %.1f°C
      Humidity    : %d%%
      Wind        : %.1f km/h
	`,
		locname,
		weather.Temperature,
		weather.Humidity,
		weather.WindSpeed,
	)

	fmt.Println("")


}
