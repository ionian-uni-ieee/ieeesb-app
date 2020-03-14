import { Sidebar } from 'components/Sidebar'
import React from 'react'
import { Route, Switch, useLocation } from 'react-router-dom'
import { CSSTransition, TransitionGroup } from 'react-transition-group'
import appVersion from 'utils/appVersion'

import Dashboard from './Dashboard'
import Events from './Events'
import Finance from './Finance'
import Settings from './Settings'
import Sponsors from './Sponsors'
import Users from './Users'

const ControlPanel: React.FC = () => {
  const location = useLocation()

  return (
    <div className='page control-panel'>
      <Sidebar>
        <div>
          <Sidebar.Button to='dashboard'>
            Dashboard
          </Sidebar.Button>
          <Sidebar.Button to='events'>
            Events
          </Sidebar.Button>
          <Sidebar.Button to='users'>
            Users
          </Sidebar.Button>
          <Sidebar.Button to='sponsors'>
            Sponsors
          </Sidebar.Button>
          <Sidebar.Button to='finance'>
            Finance
          </Sidebar.Button>
          <Sidebar.Button to='settings'>
            Settings
          </Sidebar.Button>
        </div>
        <Sidebar.Button
          id='sidebar-logout-button'
          onClick={() => null}
        >
          Logout
        </Sidebar.Button>
        <Sidebar.Footer
          text={`
            COPYRIGHTS © IONIO IEEE SB
            2020-2021
            ${appVersion}
          `}
        />
      </Sidebar>
      <TransitionGroup style={{height: '100%'}}>
        <CSSTransition
          key={location.key}
          classNames='fade'
          timeout={300}
          >
            <Switch location={location}>
              <Route
                path='/admin/dashboard'
                exact
                component={Dashboard}
              />
              <Route
                path='/admin/events'
                exact
                component={Events}
              />
              <Route
                path='/admin/users'
                exact
                component={Users}
              />
              <Route
                path='/admin/Sponsors'
                exact
                component={Sponsors}
              />
              <Route
                path='/admin/finance'
                exact
                component={Finance}
              />
              <Route
                path='/admin/settings'
                exact
                component={Settings}
              />
            </Switch>
          </CSSTransition>
      </TransitionGroup>
    </div>
  )
}

export default ControlPanel
