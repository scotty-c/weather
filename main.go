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
	Lattitude  float64 `json:"latitude"`  // Latitude of the IP address
	Longtitude float64 `json:"longitude"` // Longitude of the IP address
}

// Define a maximum number of retries to prevent infinite recursion.
const maxRetries = 5

// locationWithRetry attempts to fetch the location, retrying up to maxRetries times if the city is not found.
func locationWithRetry(retryCount int) (string, string) {
	if retryCount >= maxRetries {
		// If the maximum number of retries is reached, return an error or empty strings.
		fmt.Println("Maximum retries reached. Unable to fetch location.")
		return "", ""
	}

	url := "https://api.seeip.org/geoip"
	resp, err := http.Get(url)
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
		return locationWithRetry(retryCount + 1)
	}

	// Add a return statement at the end of the function.
	return data.City, data.Country
}

// wttr fetches and returns the weather information for the given location.
func wttr() string {
	city, country := locationWithRetry(0)                           // Get the city and country for the IP address
	u := "https://v3.wttr.in/" + city + "+" + country + "?format=3" // Construct the URL for the weather API
	resp, err := http.Get(u)                                        // Send a GET request to the weather API
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
	fmt.Println(string(wttr())) // Print the weather information for the IP address's location
}
