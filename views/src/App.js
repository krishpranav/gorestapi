import React from 'react';
import logo from './logo.svg';
import './App.css';

class RestApiFrontend extends React.Component {

  componentDidMount() {
    const apiUrl = 'http://localhost:8080/about';
    fetch(apiUrl)
      .then((response) => response.json())
      .then((data) => console.log('This is your data', data));
  }

  albumapirequest() {
    const albumrequest = 'http://localhost:8080/albums'
    fetch(albumrequest)
      .then((response) => response.json())
      .then((data) => console.log('Album Request'));
  }

  render() {

    this.albumapirequest()

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
export default RestApiFrontend;
