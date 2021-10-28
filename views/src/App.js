import React from 'react';
import logo from './logo.svg';
import './App.css';

class RestApiFrontend extends React.Component {

  componentDidMount() {
    const aboutrequest = 'http://localhost:8080/about';
    fetch(aboutrequest)
      .then((response) => response.json())
      .then((data) => console.log('About Api Request', data));
  }

  albumapirequest() {
    const albumrequest = 'http://localhost:8080/albums'
    fetch(albumrequest)
      .then((response) => response.json())
      .then((data) => console.log('Album Request', data));
  }

  artistapirequest() {
    const artistrequest = 'http://localhost:8080/artist'
    fetch(artistrequest)
      .then((response) => response.json())
      .then((data) => console.log('Arist Request', data))
  }

  profileapirequest() {
    const profilerequest = 'http://localhost:8080/profile'
    fetch(profilerequest)
      .then((response) => response.json())
      .then((data) => console.log('Profile Api Request', data))
  }

  toptracksapirequest() {
    const toptracksrequest = 'http://localhost:8080/toptracks'
    fetch(toptracksrequest)
      .then((response) => response.json())
      .then((data) => console.log('Top Tracks Api Request', data))
  }

  prodcastapirequest() {
    const prodcastrequest = 'http://localhost:8080/prodcast'
    fetch(prodcastrequest)
      .then((response) => response.json())
      .then((data) => console.log('Prodcast Api Request', data))
  }
  
  render() {

    /** 
     * START
     * api request calling area
    */
    this.profileapirequest()
    this.albumapirequest()
    this.artistapirequest()
    this.toptracksapirequest()

    /**
     * END
     */
    
    /**
     * FRONTEND 
     */
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
