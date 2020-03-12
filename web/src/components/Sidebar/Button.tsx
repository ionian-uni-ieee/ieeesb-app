import React from 'react'
import { useHistory, useLocation } from 'react-router-dom'

interface IProps {
  className?: string
  id?: string
  to?: string
}

const Button: React.FC<IProps> = (props) => {
  const {
    children,
    className,
    id,
    to,
  } = props
  const location = useLocation()
  const { pathname } = location
  const cpPage = pathname.substr(pathname.lastIndexOf('/')+1)
  const isPathActive = cpPage === to

  const history = useHistory()
  const handleClick = () => {
    if (to && !isPathActive) {
      history.push(to)
    }
  }

  return (
    <button className={`
      sidebar__button
      ${isPathActive ? 'sidebar__button_active' : ''}
      ${className}
    `}
    onClick={handleClick}
    id={id}
    >
      {children}
    </button>
  )
}

export default Button
