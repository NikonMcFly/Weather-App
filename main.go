package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
)
 //url https://api.forecast.io/forecast/APIKEY/LATITUDE,LONGITUDE

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/weather", handleWeather)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWeather(res http.ResponseWriter, req *http.Request) {
	var f Forecast
	result, err := http.Get("https://api.forecast.io/forecast/4fa91806686db356f50aa970166d6ced/40.7142,-74.0064")
	if err != nil {
		fmt.Println("something went wrong here with url", err)
		return
	}


	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)

	err = json.Unmarshal(body, &f)

	if err != nil {
		fmt.Println("something went wrong here with json :", err)
		return
	}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
  res.WriteHeader(http.StatusOK)
  json.NewEncoder(res).Encode(f)
}