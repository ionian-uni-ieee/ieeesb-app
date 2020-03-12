import { Logo } from 'components/Logo'
import React from 'react'

import Button from './Button'
import Divider from './Divider'

interface IProps {
  className?: string
  id?: string
}

interface ISidebar extends React.FC<IProps> {
  Button: typeof Button
  Divider: typeof Divider
}

const Sidebar: ISidebar = (props) => {
  const { className, id, children } = props
  return (
    <div
      className={`
        sidebar
        ${className}
      `}
      id={id}
    >
      <Logo
        width='fit-content'
        height='38px'
        white
      />
      <Divider />
      {children}
    </div>
  )
}

Sidebar.Button = Button
Sidebar.Divider = Divider 

export default Sidebar
