package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Define the structure to match the JSON response from the API
type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// Function to fetch weather data
func fetchWeatherData(city, apiKey string) (*WeatherData, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}

// Function to display weather data
func displayWeatherData(weatherData *WeatherData) {
	fmt.Printf("City: %s\n", weatherData.Name)
	fmt.Printf("Temperature: %.2f°C\n", weatherData.Main.Temp)
	fmt.Printf("Weather: %s\n", weatherData.Weather[0].Description)
}

func main() {
	apiKey := os.Getenv("6ac05326a40e454c7474bf429b9243f2")
	if apiKey == "" {
		log.Fatalf("API key not set")
	}

	fmt.Print("Enter city name: ")
	var city string
	fmt.Scanln(&city)

	weatherData, err := fetchWeatherData(city, apiKey)
	if err != nil {
		log.Fatalf("Failed to fetch weather data: %v", err)
	}

	displayWeatherData(weatherData)
}
