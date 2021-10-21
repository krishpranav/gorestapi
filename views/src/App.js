import React from 'react';
import logo from './logo.svg'
import './App.css'

class ApiRequest extends React.Component {
  albumrequest() {
    const album = 'http://localhost:8080/albums';
    fetch(album)
      .then((response) => response.json())
      .then((data) => console.log('Album Details', data));
  }

  profilerequest() {
      const profile = 'http://localhost:8080/profile';
      fetch(profile)
          .then((response) => response.json())
          .then((data) => console.log('Profile Details', data));
  }

  aboutrequest() {
      const aboutapi = 'http://localhost:8080/about';
      fetch(aboutapi)
          .then((response) => response.json())
          .then((data) => console.log('About Api Details', data));
  }

  profilerequest() {
    const profileapi = 'http://localhost:8080/profile'
    fetch(profileapi)
      .then((response) => response.json())
      .then((data) => console.log('Profile Request Details', data));
  }

    render() {
    return(
        this.profilerequest(),
        this.albumrequest(),
        this.aboutrequest(),
        this.profilerequest(),
      <div className="App">
        <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        
        <h2>Check Out The Console For Album Details</h2>
        </header>
      </div>
    )
  }
}


export default ApiRequest;