import React from 'react'
import { Calendar } from 'components/Calendar'

const Events: React.FC = () => (
  <>
    <Calendar>
      <div className="card">
        <Calendar.Eventcard>
          <Calendar.Eventcard.Date
            day={12}
            month='Jun'
            weekday='Mon'
          />
          <Calendar.Eventcard.Description
            id='1241241'
            name='Arduino Workshop'
            type='Workshop'
            tags={['Programming', 'Beginner', 'Workshop', 'Denexwidea']}
            time='12:00'
          />
          <Calendar.Eventcard.Description
            id='asdgasdagsa'
            name='Learn C++'
            type='Seminar'
            tags={['WTFLanguage', 'NotBeginner', 'Seminar', 'WorstProgrammingLanguage']}
            time='12:00'
          />
        </Calendar.Eventcard>

        <Calendar.Eventcard>
          <Calendar.Eventcard.Date
            day={12}
            month='Jun'
            weekday='Mon'
          />
          <Calendar.Eventcard.Description
            id='1241241'
            name='Arduino Workshop'
            type='Workshop'
            tags={['Programming', 'Beginner', 'Workshop', 'Denexwidea']}
            time='12:00'
          />
        </Calendar.Eventcard>
        <Calendar.Eventcard>
          <Calendar.Eventcard.Date
            day={12}
            month='Jun'
            weekday='Mon'
          />
          <Calendar.Eventcard.Description
            id='1241241'
            name='Arduino Workshop'
            type='Workshop'
            tags={['Programming', 'Beginner', 'Workshop', 'Denexwidea']}
            time='12:00'
          />
          <Calendar.Eventcard.Description
            id='asdgasdagsa'
            name='Learn C++'
            type='Seminar'
            tags={['WTFLanguage', 'NotBeginner', 'Seminar', 'WorstProgrammingLanguage']}
            time='12:00'
          />
        </Calendar.Eventcard>

      </div>
    </Calendar>
  </>
)


export default Events
