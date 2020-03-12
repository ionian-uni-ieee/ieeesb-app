import React from 'react'

interface IProps {
  active?: boolean
  className?: string
  id?: string
}

const Button: React.FC<IProps> = (props) => {
  const { children, active, className, id } = props
  return (
    <button className={`
      sidebar__button
      ${active ? 'sidebar__button_active' : ''}
      ${className}
    `}
      id={id}
    >
      {children}
    </button>
  )
}

export default Button
