import React from 'react'

interface IProps {
  active?: boolean
}

const Button: React.FC<IProps> = (props) => {
  const { children, active } = props
  return (
    <button className={`
      sidebar__button
      ${active ? 'sidebar__button_active' : ''}
    `}
    >
      {children}
    </button>
  )
}

export default Button
