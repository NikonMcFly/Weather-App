import React, { Component } from 'react';
import actions from '../redux/actions';

class List extends Component {
  render(){
    return (
    <ul>
      {
        this.props.weather.map((data) => {
          return <li key={data.id}>{data.Time}</li>
        })
      }
    </ul>
    )
  }
}

export default List
