import React from 'react'

type weekday = 'Mon' | 'Tue' | 'Wed' | 'Thu' | 'Fri' | 'Sat' | 'Sun'
type month = 'Jan' | 'Feb' | 'Mar' | 'Apr' | 'May' | 'Jun' | 'Jul' | 'Aug' | 'Sep' | 'Oct' | 'Nov' | 'Dec'

interface IDateProps {
  day: number
  weekday: weekday
  month: month
}

const Date: React.FC<IDateProps> = props => {
  const {
    day,
    weekday,
    month
  } = props

  return (
    <div className='calendar__card__date'>
      <div className="month-weekday">
        <span className='calendar__card__date__month'>
          {month}
        </span>
        <span className='calendar__card__date__weekday'>
          {weekday}
        </span>
      </div>
      <span className='calendar__card__date__day'>
        {day}
      </span>
    </div>
  )
}

export default Date
