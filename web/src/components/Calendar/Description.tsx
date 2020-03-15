import React, { ReactElement } from 'react'
import Divider from './Divider'

type type = 'Seminar' | 'Workshop'

interface IProps {
  children: ReactElement | Array<ReactElement>
}

interface IEventcard extends React.FC<IProps> {
  Date: typeof Date
  Description: typeof Description
  Divider: typeof Divider
}
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
    type,
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
      <Divider />
      <span className='calendar__eventcard__description__tags'>
        {tags.map((item: string, index: number)=><p className='tag' key={index}>{item}</p>)}
      </span>
    </div>
  )
}
export default Description
