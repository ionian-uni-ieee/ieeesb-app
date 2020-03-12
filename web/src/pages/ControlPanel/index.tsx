import { Sidebar } from 'components/Sidebar'
import React from 'react'
import { useLocation } from 'react-router-dom'

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
        >
          Logout
        </Sidebar.Button>
        <Sidebar.Footer
          text={`
            COPYRIGHTS © IONIO IEEE SB
            2020-2021
            v1.0
          `}
        />
      </Sidebar>
    </div>
  )
}

export default ControlPanel
