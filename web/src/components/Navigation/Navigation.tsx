import React, { ReactElement } from 'react'

import Button from './Button'

interface IProps {
  children: ReactElement | Array<ReactElement>
}

interface INavigation extends React.FC<IProps> {
  Button: typeof Button
}

const Navigation: INavigation = (props) => {
  return (
    <button className='ui navigation'>
      {props.children}
    </button>
  )
}

Navigation.Button = Button

export default Navigation
