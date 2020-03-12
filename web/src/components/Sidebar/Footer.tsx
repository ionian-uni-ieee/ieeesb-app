import React from 'react'

import Divider from './Divider'

interface IProps {
  text?: string
}

const Footer: React.FC<IProps> = (props) => {
  const { text, children } = props

  let textWithLineBreaks = null
  if (text) {
    const splitText = text.split(/\n/g)
    textWithLineBreaks = splitText.map(
      (subtext, index) => {
        const isLastText = index === (splitText.length - 1)
        if (isLastText) {
          return subtext
        }
        return (
          <>
            {subtext.trim()}
            <br />
          </>
        )
      }
    )
  }

  return (
    <div className='sidebar__footer'>
      <Divider />
      {!!text &&
        <p className='sidebar__footer__text'>
          {textWithLineBreaks}
        </p>
      }
      {children}
    </div>
  )
}

export default Footer
