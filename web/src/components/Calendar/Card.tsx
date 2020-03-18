import React from 'react'

import IDateProps from './Date'
import IEventProps from './Event'

interface ICard {
  date: typeof IDateProps
  events: typeof IEventProps[]
}

const Card: React.FC<ICard> = props => {
  const {
    date,
    events
  } = props
  return (
    <div className='calendar__Card'>
      {date}
      {events}
    </div>
    )

}


export default Card
