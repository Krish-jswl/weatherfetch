package weatherapi

import (
	"encoding/json"
	"net/http"
)

type WeatherResponse struct {
	Current CurrentWeather `json:"current"`
}

type CurrentWeather struct {
	Time        string  `json:"time"`
	Temperature float64 `json:"temperature_2m"`
	Humidity    int     `json:"relative_humidity_2m"`
	WindSpeed   float64 `json:"wind_speed_10m"`
}

func GetCurrentWeather(lat, lon string) (CurrentWeather, error) {

	url := "https://api.open-meteo.com/v1/forecast?latitude=" + lat + "&longitude=" + lon + "&current=temperature_2m,relative_humidity_2m,wind_speed_10m,weather_code&timezone=auto"
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		print(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	var wr WeatherResponse

	if err := json.NewDecoder(resp.Body).Decode(&wr); err != nil {
		return CurrentWeather{}, err
	}

	return wr.Current, nil
}

