package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Response struct {
	IP         string  `json:"ip"`
	City       string  `json:"city"`
	Region     string  `json:"region"`
	Country    string  `json:"country"`
	Lattitude  float64 `json:"latitude"`
	Longtitude float64 `json:"longitude"`
}

func location() (string, string) {
	//url := "https://ipapi.co/json/"
	url := "https://api.seeip.org/geoip"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if data.City == "" {
		return location()
	}

	return data.City, data.Country
}

func wttr() string {
	city, country := location()
	u := "https://v3.wttr.in/" + city + "+" + country + "?format=3"
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	bodyString := string(body)
	bodyString = strings.ReplaceAll(bodyString, "+", "")
	bodyString = strings.ReplaceAll(bodyString, country, "")

	return bodyString
}

func main() {
	fmt.Println(string(wttr()))
}
