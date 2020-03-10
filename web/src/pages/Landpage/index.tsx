import React from 'react'

import { ReactComponent as Logo } from '../../assets/svg/logo.svg'
import Navigation from '../../components/Navigation'

const Landpage = () => {
    return (
        <div className='page landpage'>
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
            <Logo
              className='page landpage__logo'
              width='80%'
              viewBox='0 0 914 231'
              style={{
                height: '400px',
                padding: '80px',
              }}
            />
        </div>
    )
}


export default Landpage
