package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Configuration struct {
	APIKey string `json:"APIKey"`
}

type WeatherInfo struct {
	CityName string `json:"name"`
	Main     struct {
		Temperature float64 `json:"temp"`
		Pressure    float64 `json:"pressure"`
		Humidity    float64 `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

type DisplayWeatherInfo struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Pressure    float64 `json:"pressure"`
	Humidity    float64 `json:"humidity"`
	Description string  `json:"description"`
}

type TemperatureUnit int

const (
	Celsius TemperatureUnit = iota
	Fahrenheit
)

func loadConfig(fileName string) (Configuration, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return Configuration{}, err
	}

	var config Configuration
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return Configuration{}, err
	}
	return config, nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Greetings from a Go server!\n"))
}

func fetchWeather(city string, unit TemperatureUnit) (DisplayWeatherInfo, error) {
	config, err := loadConfig(".apiConfig")
	if err != nil {
		return DisplayWeatherInfo{}, err
	}

	unitStr := "metric" // Default to Celsius
	if unit == Fahrenheit {
		unitStr = "imperial"
	}

	apiURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s", city, config.APIKey, unitStr)
	resp, err := http.Get(apiURL)
	if err != nil {
		return DisplayWeatherInfo{}, err
	}

	defer resp.Body.Close()

	var data WeatherInfo

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return DisplayWeatherInfo{}, err
	}

	displayInfo := DisplayWeatherInfo{
		City:        data.CityName,
		Temperature: data.Main.Temperature,
		Pressure:    data.Main.Pressure,
		Humidity:    data.Main.Humidity,
		Description: data.Weather[0].Description,
	}

	return displayInfo, nil
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			unitParam := r.URL.Query().Get("unit")
			unit := Celsius // Default to Celsius

			if unitParam == "fahrenheit" {
				unit = Fahrenheit
			}

			info, err := fetchWeather(city, unit)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(info)
		},
	)
	http.ListenAndServe(":8580", nil)
}
