package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)
 //url https://api.forecast.io/forecast/APIKEY/LATITUDE,LONGITUDE

func main() {

	http.HandleFunc("/weather", handleWeather)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
/* Something is wrong with parsing json.
Error : There was an error parsing the JSON document. 
				The document may not be well-formed.
	
*/
func handleWeather(res http.ResponseWriter, req *http.Request) {
	// var f Forecast
	result, err := http.Get("https://api.forecast.io/forecast/4fa91806686db356f50aa970166d6ced/40.7142,-74.0064")
	if err != nil {
		fmt.Println("something went wrong here with url", err)
		return
	}


	defer result.Body.Close()
	var Forecast struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timezone  string    `json:"timezone"`
	Offset    float64   `json:"offset"`
	Currently struct {
		Time        int64
		Summary     string
		Temperature float32
		Humidity    float32
		WindSpeed   float32
		WindBearing float32
	} `json="currently"`
	Daily struct {
		Summary string
		Data    []struct {
			Time           int64
			Summary        string
			TemperatureMin float32
			TemperatureMax float32
			Humidity       float32
			WindSpeed      float32
			WindBearing    float32
		} `json="data"`
	} `json="daily"`
}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	err = json.NewDecoder(result.Body).Decode(&Forecast)


	if err != nil {
		fmt.Println("something went wrong here with json :", err)
		return
	}
	
	fmt.Fprintln(res, Forecast)
}