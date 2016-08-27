import React from "react";
import ReactDOM from "react-dom";
import  { Provider } from 'react-redux';
import App from "./components/App";
import configureStore from './redux/store';

let store = configureStore();

const app = document.getElementById('app');
ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>,
    app);
