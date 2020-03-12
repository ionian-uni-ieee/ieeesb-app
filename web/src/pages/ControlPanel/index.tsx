import { Sidebar } from 'components/Sidebar'
import React from 'react'
import { useLocation } from 'react-router-dom'

const ControlPanel = () => {
  const location = useLocation()
  const { pathname } = location
  const cpPage = pathname.substr(pathname.lastIndexOf('/')+1)
  console.log(cpPage)

  return (
    <div className='page control-panel'>
      <Sidebar>
        <div>
          <Sidebar.Button
            active={cpPage === 'dashboard'}
          >
            Dashboard
          </Sidebar.Button>
          <Sidebar.Button
            active={cpPage === 'events'}
          >
            Events
          </Sidebar.Button>
          <Sidebar.Button
            active={cpPage === 'users'}
          >
            Users
          </Sidebar.Button>
          <Sidebar.Button
            active={cpPage === 'sponsors'}
          >
            Sponsors
          </Sidebar.Button>
          <Sidebar.Button
            active={cpPage === 'finance'}
          >
            Finance
          </Sidebar.Button>
          <Sidebar.Button
            active={cpPage === 'settings'}
          >
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
