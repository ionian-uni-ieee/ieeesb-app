import React from 'react'

interface IButtonProps {
  label: string
  href: string
}

const Button: React.FC<IButtonProps> = (props) => {
  const {
    label,
    href,
  } = props

  return (
      <a
        href={href}
        className='ui navigation__button'
      >
        {label}
      </a>
  )
}

export default Button