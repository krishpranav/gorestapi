/* imports */
import React from 'react';
import logo from './logo.svg'
import './App.css'

/* api request class */
class ApiRequest extends React.Component {

  /* album api request */
  albumrequest() {
    const album = 'http://localhost:8080/albums';
    fetch(album)
      .then((response) => response.json())
      .then((data) => console.log('Album Details', data));
  }

  /* profile api request */
  profilerequest() {
      const profile = 'http://localhost:8080/profile';
      fetch(profile)
          .then((response) => response.json())
          .then((data) => console.log('Profile Details', data));
  }
  
  /* about api request */
  aboutrequest() {
      const aboutapi = 'http://localhost:8080/about';
      fetch(aboutapi)
          .then((response) => response.json())
          .then((data) => console.log('About Api Details', data));
  }

  /* recommende album api request */
  recommendealbumrequest() {
    const recommendedapi = 'http://localhost:8080/toptracks'
    fetch(recommendedapi)
      .then((response) => response.json())
      .then((data) => console.log('Recommended Album Details', data))
  }

  /* prodcast api request */
  prodcastrequest() {
    const prodcastapi = 'http://localhost:8080/prodcast'
    fetch(prodcastapi)
      .then((response) => response.json())
      .then((data) => console.log('Prodcast Album Details', data))
  }
  
    render() {
    return(
        /* call back the api request functions */
        this.albumrequest(),
        this.aboutrequest(),
        this.profilerequest(),
        this.recommendealbumrequest(),
        this.prodcastrequest(),

      <div className="App">
        <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <h2>The Api Response Will Be In The Console</h2>
        </header>
      </div>
    )
  }
}

/* export the apirequest class to run it :) */
export default ApiRequest;