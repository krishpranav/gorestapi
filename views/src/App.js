import logo from './logo.svg';
import './App.css';
import React from 'react';

// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           need to develop the frontend.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           check out the github link for gorestapi
//         </a>
//       </header>
//     </div>
//   );
// }

class myComponent extends React.Component {
  componentDidMount() {
    const apiUrl = 'http://localhost:8080/albums';
    fetch(apiUrl)
      .then((response) => response.json())
      .then((data) => console.log('This is your data', data));
  }

  render() {
    return <h1> my Component has mounted, check the browser 'console' </h1>;
  }
}

export default myComponent;

// export default App;