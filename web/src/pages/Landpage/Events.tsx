import React from 'react'
import { Calendar } from 'components/Calendar'

const Events: React.FC = () => (
  <>
    <Calendar>
      <div className="card">
        <Calendar.Eventcard>
          <Calendar.Eventcard.Date
            day={19}
            month='Jun'
            weekday='Tue'
          />
          <Calendar.Eventcard.Description
            id='1241241'
            name='Arduino For Beginners'
            type='Seminar'
            tags={['Programming', 'Hardware', 'Arduino', 'Beginners']}
            time='14:00'
          />
          <Calendar.Eventcard.Description
            id='asdgasdagsa'
            name='C++ MASTERY - LEARNING HOW TO CODE'
            type='Seminar'
            tags={['programming', 'cpp']}
            time='13:00'
          />
        </Calendar.Eventcard>

        <Calendar.Eventcard>
          <Calendar.Eventcard.Date
            day={24}
            month='Jun'
            weekday='Sat'
          />
          <Calendar.Eventcard.Description
            id='1241241'
            name='Fullstack Web development - ReactJS + NodeJS'
            type='Seminar'
            tags={['Programming', 'Webdev', 'Beginners']}
            time='19:00'
          />
        </Calendar.Eventcard>
        <Calendar.Eventcard>
          <Calendar.Eventcard.Date
            day={28}
            month='Jun'
            weekday='Wed'
          />
          <Calendar.Eventcard.Description
            id='1241241'
            name='C++ MASTERY - BECOMING A PRO'
            type='Seminar'
            tags={['Programming', 'Cpp']}
            time='11:00'
          />
        </Calendar.Eventcard>

      </div>
    </Calendar>
  </>
)


export default Events
