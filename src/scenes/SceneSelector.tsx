import React, { Component } from 'react';
import { Landpage } from './Landpage';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import './style.scss';

export class SceneSelector extends Component {
  render() {
    return (
      <div className='scene'>
        <Landpage />
      </div>
    );
  }
}
