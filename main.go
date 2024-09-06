package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Define a struct to model the JSON response from the geo-location API.
type Response struct {
	IP      string `json:"ip"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// locationWithRetry attempts to fetch the location, retrying up to maxRetries times if the city is not found.
func locationWithRetry(apiURL string, retryCount int) (string, string) {
	const maxRetries = 5
	if retryCount >= maxRetries {
		fmt.Println("Maximum retries reached. Unable to fetch location.")
		return "", ""
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", ""
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", ""
	}

	if data.City == "" {
		return locationWithRetry(apiURL, retryCount+1)
	}

	return data.City, data.Country
}

// wttr fetches and returns the weather information for the given location.
func wttr(geoAPIURL, weatherAPIBaseURL string) string {
	city, country := locationWithRetry(geoAPIURL, 0)
	if city == "" {
		return "Could not determine location."
	}

	// Combine city and country to create a specific query, e.g., "Newcastle,Australia" or "Newcastle,UK"
	location := fmt.Sprintf("%s,%s", url.QueryEscape(city), url.QueryEscape(country))

	// Construct the weather API URL to return both weather emoji and current temperature.
	u := weatherAPIBaseURL + location + "?format=%c+%t"
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println("Error: ", err)
		return "Unable to fetch weather."
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return "Unable to read weather data."
	}

	// Combine the city name with the weather data
	return fmt.Sprintf("%s: %s", city, string(body))
}

func main() {
	geoAPIURL := "https://api.seeip.org/geoip"
	weatherAPIBaseURL := "https://v3.wttr.in/"
	fmt.Println(wttr(geoAPIURL, weatherAPIBaseURL)) // Print the city name along with current weather condition and temperature
}
