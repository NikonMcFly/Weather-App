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
/* Something is wrong with parsing json.
Error : There was an error parsing the JSON document. 
				The document may not be well-formed.
	
*/
func handleWeather(res http.ResponseWriter, req *http.Request) {
	var f Forecast
	result, err := http.Get("https://api.forecast.io/forecast/d671b89e0e06f52b79f833c81d43cc55/40.7142,-74.0064")
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}


	defer result.Body.Close()

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	err = json.NewDecoder(result.Body).Decode(&f)

	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
}