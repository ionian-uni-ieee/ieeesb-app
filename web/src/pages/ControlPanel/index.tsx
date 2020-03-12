import { Sidebar } from 'components/Sidebar'
import React from 'react'
import appVersion from 'utils/appVersion'

const ControlPanel = () => {
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
    </div>
  )
}

export default ControlPanel
