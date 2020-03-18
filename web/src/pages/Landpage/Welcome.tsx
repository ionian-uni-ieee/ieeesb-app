import { ReactComponent as Logo } from 'assets/svg/logo.svg'
import React from 'react'
import { scrollToID } from 'utils/browserHelpers'

const Welcome: React.FC = () => {
  const scrollToWelcome = (e: React.MouseEvent): void => {
    e.preventDefault()
    scrollToID('#welcome')
  }

  return (
    <div className='landpage__section landpage__section__welcome' id='welcome'>
      <a href='#welcome' onClick={scrollToWelcome}>
        <Logo
          width='80%'
          viewBox='0 0 914 231'
          style={{
            height: '14vw',
            padding: '14vh 0',
          }}
        />
      </a>
    </div>
  )
}

export default Welcome
