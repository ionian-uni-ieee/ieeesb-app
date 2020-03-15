import React from 'react'

interface IHeaderProps {
  name: string
}

const Header: React.FC<IHeaderProps> = props => {
  const {
    name
  } = props

  return (
    <h4 className="calendar__header">
      {name}
    </h4>
  )
}

export default Header
