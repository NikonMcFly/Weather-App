

let initial = {
		weather: {},
		daily: {},
		fetched: false,
}
let weatherReducer = function(state = initial, action){
  switch (action.type){
  	case "RECEIVE_WEATHER":
			return {
					weather: action.data,
					fetched: true,
				};
  	default:
  		return state;
  }
}

export default weatherReducer
