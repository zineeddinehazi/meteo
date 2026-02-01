package main

import (
	"encoding/json"
	"example/meteo/pkg"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	API_KEY := os.Getenv("API_KEY")
	city := pkg.GetUserInput()
	encodedCity := url.QueryEscape(city)
	numOfDays := 3

	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=%d",
		API_KEY, encodedCity, numOfDays)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status: %s", res.Status)
	}

	var data pkg.ForecastResponse

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	pkg.DisplayTable(data, city)
}
