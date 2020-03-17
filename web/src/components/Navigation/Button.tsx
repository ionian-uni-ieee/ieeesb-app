import React from 'react'
import { useHistory } from 'react-router'

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
    const regexIsID = /^\/?#[a-zA-Z\d-]+/
    const isLinkID = regexIsID.test(href)
    if (!isLinkID) {
      return
    }
    const id = href.startsWith('/') ? href.substr(2) : href.substr(1)
    const idElement = document.getElementById(id)
    if (!idElement) {
      return
    }
    const { offsetTop } = idElement
    window.scroll({
      top: offsetTop,
      left: 0,
      behavior: 'smooth'
    })
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
