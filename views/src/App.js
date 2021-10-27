import React from 'react';
import logo from './logo.svg';
import './App.css';

class myComponent extends React.Component {
  componentDidMount() {
    const apiUrl = 'http://localhost:8080/about';
    fetch(apiUrl)
      .then((response) => response.json())
      .then((data) => console.log('This is your data', data));
  }

  render() {

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          View The Console For Api Request
        </p>
      </header>
    </div>
  );

  }
}
export default myComponent;
