import { Logo } from 'components/Logo'
import React from 'react'

import Button from './Button'
import Divider from './Divider'
import Footer from './Footer'

interface IProps {
  className?: string
  id?: string
}

interface ISidebar extends React.FC<IProps> {
  Button: typeof Button
  Divider: typeof Divider
  Footer: typeof Footer
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
Sidebar.Footer = Footer

export default Sidebar
