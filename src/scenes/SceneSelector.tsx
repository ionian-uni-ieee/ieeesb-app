import React, { Component } from 'react';
import { Landpage } from './Landpage';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import './style.scss';

export class SceneSelector extends Component {
  render() {
    return (
      <div className='scene'>
<<<<<<< HEAD
          <Router>
              <div className='routes'>
                <Route path="/" exact={true} component={Landpage} />
              </div>
          </Router>
=======
        <Landpage />
>>>>>>> 18a116f019f4e1d3c070c93e423c2dc4ce569802
      </div>
    );
  }
}
