package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Define a struct to model the JSON response from the geo-location API.
type Response struct {
	IP         string  `json:"ip"`        // IP address of the requestor
	City       string  `json:"city"`      // City of the IP address
	Region     string  `json:"region"`    // Region or state of the IP address
	Country    string  `json:"country"`   // Country of the IP address
	Latitude   float64 `json:"latitude"`  // Latitude of the IP address
	Longitude  float64 `json:"longitude"` // Longitude of the IP address
}

// Define a maximum number of retries to prevent infinite recursion.
const maxRetries = 5

// locationWithRetry attempts to fetch the location, retrying up to maxRetries times if the city is not found.
func locationWithRetry(apiURL string, retryCount int) (string, string) {
	if retryCount >= maxRetries {
		// If the maximum number of retries is reached, return an error or empty strings.
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
		// Retry with incremented retry count if city is not found.
		return locationWithRetry(apiURL, retryCount+1)
	}

	return data.City, data.Country
}

// wttr fetches and returns the weather information for the given location.
func wttr(geoAPIURL, weatherAPIBaseURL string) string {
	city, country := locationWithRetry(geoAPIURL, 0)                           // Get the city and country for the IP address
	u := weatherAPIBaseURL + city + "+" + country + "?format=3" // Construct the URL for the weather API
	resp, err := http.Get(u)                                    // Send a GET request to the weather API
	if err != nil {
		fmt.Println("Error: ", err) // Print error if the request fails
	}
	defer resp.Body.Close() // Ensure the response body is closed after processing

	body, err := io.ReadAll(resp.Body) // Read the entire response body
	if err != nil {
		fmt.Println("Error: ", err) // Print error if reading the body fails
	}

	bodyString := string(body)                               // Convert the response body to a string
	bodyString = strings.ReplaceAll(bodyString, "+", "")     // Remove '+' characters from the response
	bodyString = strings.ReplaceAll(bodyString, country, "") // Remove the country name from the response

	return bodyString // Return the weather information as a string
}

func main() {
	geoAPIURL := "https://api.seeip.org/geoip"
	weatherAPIBaseURL := "https://v3.wttr.in/"
	fmt.Println(wttr(geoAPIURL, weatherAPIBaseURL)) // Print the weather information for the IP address's location
}
