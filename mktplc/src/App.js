import React, { Component } from 'react';
import {  Jumbotron } from 'react-bootstrap';
import logo from './logo.png';
import History from './components/details';

//import Artland from './Artland'

import './App.css';


class App extends Component {
  render() {
    return (
      <div>
        <Jumbotron>
        <History route="get"></History>
        </Jumbotron>
      </div>
    );
  }
}


export default App;
