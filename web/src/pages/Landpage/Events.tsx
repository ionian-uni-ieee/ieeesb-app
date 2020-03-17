import React from 'react'
import { Calendar } from 'components/Calendar'

const Events: React.FC = () => (
  <div className='landpage__section landpage__section__events'>
    <Calendar>
      <div className="card-item">
        {api.getEvents().map(event => (<Calendar.Card Events={} Date={} />))}
      </div>
    </Calendar>
  </div>
)


export default Events
