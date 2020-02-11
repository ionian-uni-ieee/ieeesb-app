import React from 'react'

import Navigation from '../../components/Navigation'

const Landpage = () => {
    return (
        <div className='pages landpage'>
            <Navigation>
              <Navigation.Button
                label='Events'
                href='#events'
              />
              <Navigation.Button
                label='About us'
                href='#aboutus'
              />
            </Navigation>
        </div>
    )
}


export default Landpage 
