package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"

)
 //url https://api.forecast.io/forecast/APIKEY/LATITUDE,LONGITUDE
//here is an example https://github.com/kpurdon/go-api-example


type Forecast struct {
		Latitude 	float64 
		Longitude float64
		Timezone 	string
	Currently struct {
		Time        int64
		Summary     string
		Temperature float32
		Humidity    float32
		WindSpeed   float32
		WindBearing float32
	} 
	Daily struct {

		Data    []struct {
			Time           int64
			Summary        string
			TemperatureMin float32
			TemperatureMax float32
			Humidity       float32
			WindSpeed      float32
			WindBearing    float32
		} 
	} 
}

//add function to change location for a given location

//func Get(){}


func handleWeather(res http.ResponseWriter, req *http.Request) {
	var f Forecast
	//find the largest cities in the US and for loop them for location
	//http://www.travelmath.com/cities/
	key := "add key"
	NYCcoord  		:= "40.7142,-74.0064"

	url := "https://api.forecast.io/forecast" + "/" + key + "/" + NYCcoord
	result, err := http.Get(url)
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
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
  res.WriteHeader(http.StatusOK)
  json.NewEncoder(res).Encode(f)
}


func main() {
	// http.Handle("/", http.FileServer(http.Dir("client")))
	http.HandleFunc("/weather", handleWeather)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
