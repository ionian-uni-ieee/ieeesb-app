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
              style={{
                height: `${window.innerHeight}px`,
                paddingTop: `${window.innerHeight / 4}px`,
              }}
            />
        </div>
    )
}


export default Landpage
