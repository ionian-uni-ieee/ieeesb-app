import React from 'react'

type type = 'Seminar' | 'Workshop'

interface IPropsDescription {
  id: string
  name: string
  tags: string[]
  time: string
  type: type
}

const Description: React.FC<IPropsDescription> = props => {
  const {
    id,
    name,
    tags,
    time,
    type
  } = props

  return (
    <div
      className='calendar__eventcard__description'
      id={id}
    >
      <p className='calendar__eventcard__description__type'>
        {type}
      </p>
      <p className='calendar__eventcard__description__name'>
        {name}
      </p>
      <p className='calendar__eventcard__description__time'>
        {time}
      </p>
      <p className='calendar__eventcard__description__tags'>
        {tags}
      </p>
    </div>
  )
}

export default Description