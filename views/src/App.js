import React from 'react';
import logo from './logo.svg'
import './App.css'

class myComponent extends React.Component {
  componentDidMount() {
    const apiUrl = 'http://localhost:8080/albums';
    fetch(apiUrl)
      .then((response) => response.json())
      .then((data) => console.log('Album Details', data));
  }
    render() {
    return(
      <div className="App">
        <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        
        <h2>Check Out The Console For Album Details</h2>
        </header>
      </div>
    )
  }
}

export default myComponent;