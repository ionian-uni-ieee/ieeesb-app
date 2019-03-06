import React, { Component } from 'react';
import { Landpage } from './Landpage';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import './style.scss';

export class SceneSelector extends Component {
  render() {
    return (
      <div className='scene'>
<<<<<<< HEAD
        <Landpage />
=======
          <Router>
              <div className='routes'>
                <Route path="/ieee" exact={true} component={Landpage} />
              </div>
          </Router>
>>>>>>> 54f84db3a4fcff24d5bc7bd636ac501f7cde0ad2
      </div>
    );
  }
}
