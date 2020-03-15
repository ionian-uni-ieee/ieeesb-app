import React from 'react'

type weekday = 'Mon' | 'Tue' | 'Wed' | 'Thu' | 'Fri' | 'Sat' | 'Sun'

interface IDateProps {
  day: number
  weekday: weekday
  month: string
}

const Date: React.FC<IDateProps> = props => {
  const {
    day,
    weekday,
    month
  } = props

  return (
    <div className='calendar__eventcard__date'>
      <div className="month-weekday">
        <span className='calendar__eventcard__date__month'>
          {month}
        </span>
        <span className='calendar__eventcard__date__weekday'>
          {weekday}
        </span>
      </div>

      <span className='calendar__eventcard__date__day'>
        {day}
      </span>

    </div>
  )
}

export default Date
