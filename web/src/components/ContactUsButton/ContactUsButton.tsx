import React from 'react'

interface IContactUsButton {
  label: string
  href: string

}

const ContactUsButton: React.FC<IContactUsButton> = props => {
  const {
    label,
    href,
  } = props

  return (
    <a 
      className = 'contact_us_button'
      href = {href}
    >
      {label}
      
    </a>
  )
}

export default ContactUsButton