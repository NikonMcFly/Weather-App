import axios from "axios";

const fetchWeather = function fetchWeather() {
  return fetch("http://localhost:8080/weather")
}


const getWeather = function getWeather() {
  return function(dispatch) {
    return fetchWeather()
      .then(
        response => response.json(),
        error => console.log(error)
      )
      .then(
        json => dispatch(receiveWeather(json))
      )
  };
}

const receiveWeather = function receiveWeather(data) {
  return {
    type: "RECEIVE_WEATHER",
    data
  }
}


// const getWeather = function getWeather(){
//   return{
//     type: "GET_WEATHER",
//     var Data =  axios.get("localhost:8080/weather")
//   }
// }


export default {  
        fetchWeather,
        getWeather
}