import React, { ReactElement } from 'react'

import Eventcard from './Eventcard'

interface IProps {
  children: ReactElement | Array<ReactElement>
}

interface ICalendar extends React.FC<IProps> {
  Eventcard: typeof Eventcard
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

Calendar.Eventcard = Eventcard

export default Calendar
