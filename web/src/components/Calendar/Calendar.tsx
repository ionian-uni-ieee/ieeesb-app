import React, { ReactElement } from 'react'

import Header from './Header'
import Eventcard from './Eventcard'

interface IProps {
  children: ReactElement | Array<ReactElement>
}

interface ICalendar extends React.FC<IProps> {
  Header: typeof Header
  Eventcard: typeof Eventcard
}

const Calendar: ICalendar = props => (
  <div
    className='calendar'
  >
    <Header name='Calendar' />
    {props.children}
  </div>
)

Calendar.Header = Header
Calendar.Eventcard = Eventcard

export default Calendar
