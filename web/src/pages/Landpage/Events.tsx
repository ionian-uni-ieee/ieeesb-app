import React from 'react'
import { Calendar } from 'components/Calendar'

const Events: React.FC = () => (
  <div className='landpage__section landpage__section__events'>
    <Calendar cards={api.getEvents()}/>
  </div>
)


export default Events
