import React from 'react'
import { BrowserRouter as Router, Redirect, Route, Switch } from 'react-router-dom'

import { ControlPanelLogin, Landpage } from './pages'

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
