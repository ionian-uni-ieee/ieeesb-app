import React, { Component } from 'react';
import './App.scss';
import { SceneSelector } from './scenes/SceneSelector';

class App extends Component {
  render() {
    return (
      <div className="App">
        <SceneSelector />
      </div>
    );
  }
}

export default App;
