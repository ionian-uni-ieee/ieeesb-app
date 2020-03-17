import React, { ReactElement } from 'react'

import Date from './Date'
import Event from './Event'

interface IProps {
  children: ReactElement | Array<ReactElement>
}

interface ICard extends React.FC<IProps> {
  Date: typeof Date
  Events: typeof Event[]
}

const Card: React.FC<ICard> = props => (
  <div className='calendar__Card'>
    {props.children}
  </div>
)


export default Card
