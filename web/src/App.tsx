import React from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'

import { Landpage } from './pages'

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
        </Switch>
      </Router>
    </div>
  )
}

export default App
