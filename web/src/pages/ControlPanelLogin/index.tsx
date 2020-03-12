import { Button } from 'components/Button'
import { Input } from 'components/Input'
import { Logo } from 'components/Logo'
import React from 'react'

const ControlPanelLogin = () => {
  return (
    <div className='page control-panel-login'>
      <div className='page control-panel-login__form'>
        <Logo
          noRectangle
          height='80px'
        />
        <Input
          placeholder='USERNAME'
        />
        <Input
          placeholder='PASSWORD'
          secret
        />
        <Button primary>
          LOGIN
        </Button>
        <Button secondary>
          REGISTER
        </Button>
      </div>
    </div>
  )
}

export default ControlPanelLogin
