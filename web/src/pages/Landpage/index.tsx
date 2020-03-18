import Navigation from 'components/Navigation'
import React, { useEffect } from 'react'

import AboutUs from './AboutUs'
import Welcome from './Welcome'
import Events from './Events'

const makeLogoWhiteOnMeetUs = (): void => {
  const {scrollY} = window
  const logoSvg = document.querySelector('.logo__svg') as HTMLElement
  const meetUsSection = document.getElementById('who-are-we') as HTMLDivElement
  const { offsetTop } = meetUsSection

  if (scrollY >= (offsetTop - 600)) {
    logoSvg.classList.add('logo_white')
  } else {
    logoSvg.classList.remove('logo_white')
  }
}

const Landpage: React.FC = () => {
  useEffect(() => {
    window.addEventListener('scroll', makeLogoWhiteOnMeetUs)

    return (): void => {
      window.removeEventListener('scroll', makeLogoWhiteOnMeetUs)
    }
  }, [])

  return (
    <div>
      <Navigation>
        <Navigation.Button
          label='ΠΟΙΟΙ ΕΙΜΑΣΤΕ'
          href='#who-are-we'
        />
        <Navigation.Button
          label='TI KANOYME'
          href='#what-we-do'
        />
        <Navigation.Button
          label='ΓΝΩΡΙΣΤΕ ΜΑΣ'
          href='#meet-us'
        />
      </Navigation>
      <div className='page landpage'>
        <Welcome />
        <Events />
        <AboutUs />
      </div>
    </div>
  )
}


export default Landpage
