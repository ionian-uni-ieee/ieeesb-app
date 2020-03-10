import React, { SyntheticEvent } from 'react'

interface IProps {
  primary?: boolean
  secondary?: boolean
  fluid?: boolean
  onClick?: (e: SyntheticEvent) => void
}

const Button: React.FC<IProps> = (props) => {
  const {
    primary,
    secondary,
    fluid,
    children,
    onClick,
  } = props

  return (
    <div
      className={`
        button
        ${secondary && !primary ? 'button_secondary' : ''}
        ${fluid ? 'button_fluid' : ''}
      `}
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
