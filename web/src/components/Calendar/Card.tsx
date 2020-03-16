import React, { ReactElement } from 'react'

import Date from './Date'
import Event from './Event'

interface IProps {
  children: ReactElement | Array<ReactElement>
}

interface ICard extends React.FC<IProps> {
  Date: typeof Date
  Event: typeof Event
}

const Card: ICard = props => (
  <div className='calendar__Card'>
    {props.children}
  </div>
)

Card.Date = Date
Card.Event = Event

export default Card
