import { ReactComponent as LogoSVG } from 'assets/svg/logo.svg'
import { ReactComponent as LogoWhiteSVG } from 'assets/svg/logo_white.svg'
import React from 'react'

interface IProps {
  white?: boolean
  noRectangle?: boolean
  width?: string
  height?: string
}

const Logo: React.FC<IProps> = (props) => {
  const {
    white,
    noRectangle,
    width,
    height,
  } = props

  const logoProps = {
    width,
    style: {
      height,
    },
  }

  return (
    <span
      className={`
        logo
        ${noRectangle ? 'logo_no-rectangle' : ''}
      `}
    >
      {!white &&
        <LogoSVG
          viewBox='0 0 914 231'
          {...logoProps}
        />
      }
      {white &&
        <LogoWhiteSVG
          viewBox='0 0 150 38'
          {...logoProps}
        />
      }
    </span>
  )
}

export default Logo
