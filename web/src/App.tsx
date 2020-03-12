import React from 'react'
import { BrowserRouter as Router, Redirect, Route, Switch } from 'react-router-dom'

import { ControlPanel, ControlPanelLogin, Landpage } from './pages'

const App: React.FC = () => {
  return (
    <div className='App'>
      <Router basename='/ieee'>
        <Switch>
          <Route
            path='/'
            exact
            component={Landpage}
          />
          <Route
            path='/admin'
            exact
            component={ControlPanelLogin}
          />
          <Route
            path='/admin/'
            component={ControlPanel}
          />
          <Redirect
            from='/'
            to='/'
          />
        </Switch>
      </Router>
    </div>
  )
}

export default App
