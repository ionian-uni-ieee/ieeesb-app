import React from 'react'
import { useHistory } from 'react-router'
import { scrollToID } from 'utils/browserHelpers'

interface IButtonProps {
  label: string
  href?: string
  onClick?(e: React.MouseEvent<HTMLAnchorElement>): void
}

const Button: React.FC<IButtonProps> = props => {
  const history = useHistory()
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
    history.push(href)
    e.preventDefault()
    scrollToID(href)
  }

  return (
    <a
      href={href}
      className='ui navigation__button'
      onClick={handleClick}
    >
      {label}
    </a>
  )
}

export default Button
