import React from 'react';  
import {RouteHandler, Link} from 'react-router';

class App extends React.Component {  
  render() {
    return (
      <div>
        <h1>Quiz</h1>
        <RouteHandler/>
      </div>
    );
  }
}

export default App;
