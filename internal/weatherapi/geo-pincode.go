package weatherapi

import (
	"encoding/json"
	"net/http"
	"errors"
)

type NominatimResult struct {
	Lat 		string		`json:"lat"`
	Lon			string		`json:"lon"`
	Name		string		`json:"name"`
	DisplayName	string		`json:"display_name"`
}


func GeoPincode(pincode string) (NominatimResult, error){

	url := "https://nominatim.openstreetmap.org/search?" + "postalcode=" + pincode + "&country=India&format=json&limit=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		print(err)
	}
	req.Header.Set("User-Agent", "weatherfetch/1.0 (github.com/yourname/weatherfetch)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	
	var results []NominatimResult

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return NominatimResult{}, err
	}

	if len(results) == 0 {
    return NominatimResult{}, errors.New("no results found")
	}

	return results[0], nil
}
