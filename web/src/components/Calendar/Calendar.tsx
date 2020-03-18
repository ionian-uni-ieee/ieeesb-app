import React from 'react'
import ICard from './Card'

interface IProps {
  cards: typeof ICard[]
}

const Calendar: React.FC<IProps> = props => {
  const {
    cards
  } = props

  return (
    <div className='calendar'>
      <h1 className="calendar__header">
        Calendar
      </h1>
      <div className="card-item">
        {cards.map(card => <>{card}</>)}
      </div>
    </div>
  )
}

export default Calendar