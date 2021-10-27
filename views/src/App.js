import React from 'react';

class myComponent extends React.Component {
  componentDidMount() {
    const apiUrl = 'http://localhost:8080/about';
    fetch(apiUrl)
      .then((response) => response.json())
      .then((data) => console.log('This is your data', data));
  }
  render() {
    return <h1>Check The Api Request In The Console Soon The Frontend Will Be Updated With New UI</h1>;
  }
}
export default myComponent;
