import React, { ReactElement } from 'react'

import Card from './Card'

interface IProps {
  children: ReactElement | Array<ReactElement>
}

interface ICalendar extends React.FC<IProps> {
  Card: typeof Card
}

const Calendar: ICalendar = props => (
  <div
    className='calendar'
  >
    <h1 className="calendar__header">
      Calendar
    </h1>
    {props.children}
  </div>
)

Calendar.Card = Card

export default Calendar
