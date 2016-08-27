import React, { Component } from 'react';
import actions from '../redux/actions';


class Input extends Component {

	constructor(){
		super()
		this.state = {
			input: ''
		}
	}
	handleChange(event){
		this.setState({
			input: event.target.value
		})
	}
	handleSubmit(event) {
		event.preventDefault()
		// this.props.receiveWeather(this.state.input)
		console.log(this.props.actions.getWeather())
	}

	render() {
		return (
			<div>
				<input type="text" 
					placeholder="working"
					value={this.state.input}
					onChange={this.handleChange.bind(this)}
					/>
					<button onClick={this.handleSubmit.bind(this)}>Submit</button>
					<p>{this.handleSubmit}</p>
			</div>
			)
		}
}

export default Input