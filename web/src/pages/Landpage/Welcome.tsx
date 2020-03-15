import { ReactComponent as Logo } from 'assets/svg/logo.svg'
import React from 'react'

const Welcome: React.FC = () => (
  <div className='landpage__section landpage__section__welcome'>
    <Logo
      width='80%'
      viewBox='0 0 914 231'
      style={{
        height: '14vw',
        padding: '14vh 0',
      }}
    />
  </div>
  )

export default Welcome
