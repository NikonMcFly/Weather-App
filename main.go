package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)
 //url https://api.forecast.io/forecast/APIKEY/LATITUDE,LONGITUDE

func main() {

	http.HandleFunc("/", handleWeather)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWeather(res http.ResponseWriter, req *http.Request) {
	var f Forecast
	result, err := http.Get("https://api.forecast.io/forecast/{API_KEY_HERE}/40.7142,-74.0064")
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}

	defer result.Body.Close()

	err = json.NewDecoder(result.Body).Decode(&f)

	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
  fmt.Fprintln(res, f)
}