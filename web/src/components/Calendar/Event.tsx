import React from 'react'

type type = 'Seminar' | 'Workshop'

interface IEventProps {
  id: string
  name: string
  tags: string[]
  time: string
  type: type
}

const Event: React.FC<IEventProps> = props => {
  const {
    id,
    name,
    tags,
    time,
    type,
  } = props

  return (
    <div
      className='calendar__card__event'
      id={id}
    >
      <p className='calendar__card__event__type'>
        {type}
      </p>
      <p className='calendar__card__event__name'>
        {name}
      </p>
      <p className='calendar__card__event__time'>
        {time}
      </p>
      <hr className='calendar__card__event__divider'/>
      <span className='calendar__card__event__tags'>
        {tags.map((item: string, index: number)=><p className='tag' key={index}>{item}</p>)}
      </span>
    </div>
  )
}
export default Event