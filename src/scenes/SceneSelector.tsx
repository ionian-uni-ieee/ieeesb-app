import React, { Component } from 'react';
import { Landpage } from './Landpage';
import { BrowserRouter as Router, Route } from 'react-router-dom';

export class SceneSelector extends Component {
  render() {
    return (
      <div className='scene'>
          <Router>
              <div>
                <Route path="/" exact={true} component={Landpage} />
              </div>
          </Router>
      </div>
    );
  }
}
