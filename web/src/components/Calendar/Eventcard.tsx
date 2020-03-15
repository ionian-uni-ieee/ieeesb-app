import React, { ReactElement } from 'react'

import Date from './Date'
import Description from './Description'

interface IProps {
  children: ReactElement | Array<ReactElement>
}

interface IEventcard extends React.FC<IProps> {
  Date: typeof Date
  Description: typeof Description
}

const Eventcard: IEventcard = props => (
  <div className='calendar__eventcard'>
    {props.children}
  </div>
)

Eventcard.Date = Date
Eventcard.Description = Description

export default Eventcard
