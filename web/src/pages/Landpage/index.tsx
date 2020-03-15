import Navigation from 'components/Navigation'
import React from 'react'

import AboutUs from './AboutUs'
import Welcome from './Welcome'

const Landpage: React.FC = () => (
  <div>
    <Navigation>
      <Navigation.Button
        label='Events'
        href='https://www.facebook.com/ieeeIonianUni/events/'
      />
      <Navigation.Button
        label='About us'
        href='https://www.facebook.com/pg/ieeeIonianUni/about/?ref=page_internal'
      />
    </Navigation>
    <div className='page landpage'>
      <Welcome />
      <AboutUs />
    </div>
  </div>
)


export default Landpage
