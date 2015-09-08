import React from 'react';  
import {RouteHandler, Link} from 'react-router';

class Main extends React.Component {  
  render() {
    return (
      <div>
        <h1>Quiz</h1>
        <RouteHandler/>
      </div>
    );
  }
}

export default Main;  
