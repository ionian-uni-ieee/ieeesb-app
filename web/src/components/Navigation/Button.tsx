import React from 'react'
import { Link } from 'react-router-dom'
import { scrollToID } from 'utils/browserHelpers'

interface IButtonProps {
  label: string
  href?: string
  onClick?(e: React.MouseEvent<HTMLAnchorElement>): void
}

const Button: React.FC<IButtonProps> = props => {
  const {
    label,
    href,
    onClick= null,
  } = props

  const handleClick = (e: React.MouseEvent<HTMLAnchorElement>): void => {
    if (onClick) {
      onClick(e)
    }
    if (!href) {
      return
    }
    e.preventDefault()
    scrollToID(href)
  }

  return (
    <Link
      to={href || 'ieee'}
      className='ui navigation__button'
      onClick={handleClick}
    >
      {label}
    </Link>
  )
}

export default Button
