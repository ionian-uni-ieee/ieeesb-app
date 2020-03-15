import React, { ReactElement, useEffect, useRef } from 'react'

import Button from './Button'

interface IProps {
  children: ReactElement | Array<ReactElement>
}

interface INavigation extends React.FC<IProps> {
  Button: typeof Button
}

const Navigation: INavigation = props => {
  const navigationRef = useRef<HTMLDivElement>(null)
  useEffect(() => {
    window.addEventListener('scroll', () => {
      const currentY = window.pageYOffset
      if (!navigationRef.current) {
        return
      }
      if (currentY > 60) {
        navigationRef.current.classList.add('navigation_sticky')
      } else {
        navigationRef.current.classList.remove('navigation_sticky')
      }
    })
  }, [])

  return (
    <div className='ui navigation' ref={navigationRef}>
      {props.children}
    </div>
  )
}

Navigation.Button = Button

export default Navigation
