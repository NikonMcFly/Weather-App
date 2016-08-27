import React, { Component } from 'react';
import Input from './input'
import List from './list'
import actions from '../redux/actions';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

class App extends Component {
	componentWillMount() {
    this.props.actions.getWeather();
  }


  // renderU(){
		// 		this.props.temp.weather.Daily.map(data=>{
		// 			return <li data={data} key={data.Time}>{data}</li>
		// 		})

  // 	}
	render() {
		// console.log(this.props.temp.Daily)
		// console.log(this.props.temp.weather.Daily)
		// debugger;
		return (
			<div>
				<h1>Steve</h1>
				<Input actions={this.props.actions}/>
				<p>{this.props.temp.weather.Timezone}</p>
				<p>{this.props.temp.weather.Latitude}</p>
				<ul>{this.props.temp.fetched ? 
				this.props.temp.weather.Daily.Data.map(data=>{
					return <li key={data.Time}>{data.Summary}</li>
				}) : null
			}</ul>
			</div>
			)
	}
}

function mapStateToProp(state){
	return state
}

function mapDispatchToProps(dispatch){
	return {
		actions: bindActionCreators(actions, dispatch)
	}
}

export default connect(mapStateToProp,mapDispatchToProps)(App)