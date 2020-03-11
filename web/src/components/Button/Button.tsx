import React, { SyntheticEvent } from 'react'

interface IProps {
  primary?: boolean
  secondary?: boolean
  fluid?: boolean
  onClick?: (e: SyntheticEvent) => void
  className?: string
  id?: string
}

const Button: React.FC<IProps> = (props) => {
  const {
    primary,
    secondary,
    fluid,
    children,
    onClick,
    className,
    id,
  } = props

  return (
    <div
      className={`
        button
        ${secondary && !primary ? 'button_secondary' : ''}
        ${fluid ? 'button_fluid' : ''}
        ${className}
      `}
      id={id}
    >
      <button
        className='button__element'
        onClick={onClick}
      >
        {children}
      </button>
    </div>
  )
}

export default Button
