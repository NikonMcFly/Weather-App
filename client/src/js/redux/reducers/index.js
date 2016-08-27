import  { combineReducers } from 'redux'
import weatherReducer from './WeatherReducer'


const rootReducer = combineReducers({
	temp: weatherReducer,
})

export default rootReducer