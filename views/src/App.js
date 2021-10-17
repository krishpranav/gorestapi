import React from 'react';
import  ReactDOM  from 'react-dom';

class AlbumComponenet extends React.Component{
  constructor(props){
    super(props);

    this.state={
      albums:[]
    };
  }

  componentDidMount() {
    fetch("http://localhost:8080/albums")
  }

  render() {
    return(
      <div>
        <h2>Albums Details</h2>
      </div>
    )
  }
}

export default AlbumComponenet;