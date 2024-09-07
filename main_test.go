package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Mock server response for the geo-location API
func mockGeoLocationHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		IP:         "123.456.789.012",
		City:       "Test City",
		Region:     "Test Region",
		Country:    "Test Country",
		Latitude:   12.34,
		Longitude:  56.78,
	}
	json.NewEncoder(w).Encode(response)
}

// Mock server response for the weather API
func mockWeatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test City: Clear skies"))
}

// Test for the locationWithRetry function
func TestLocationWithRetry(t *testing.T) {
	// Create a mock server for the geo-location API
	server := httptest.NewServer(http.HandlerFunc(mockGeoLocationHandler))
	defer server.Close()

	city, country := locationWithRetry(server.URL, 0)
	if city != "Test City" || country != "Test Country" {
		t.Errorf("Expected (Test City, Test Country), got (%s, %s)", city, country)
	}
}

// Test for the wttr function
func TestWttr(t *testing.T) {
	// Create mock servers for the geo-location and weather APIs
	geoLocationServer := httptest.NewServer(http.HandlerFunc(mockGeoLocationHandler))
	defer geoLocationServer.Close()

	weatherServer := httptest.NewServer(http.HandlerFunc(mockWeatherHandler))
	defer weatherServer.Close()

	weather := wttr(geoLocationServer.URL, weatherServer.URL+"/")
	expected := "Test City: Clear skies"
	if !strings.Contains(weather, expected) {
		t.Errorf("Expected %s, got %s", expected, weather)
	}
}


