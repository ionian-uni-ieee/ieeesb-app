import { ReactComponent as Logo } from 'assets/svg/logo.svg'
import { Button } from 'components/Button'
import { Input } from 'components/Input'
import React from 'react'

const ControlPanelLogin = () => {
  return (
    <div className='page control-panel-login'>
      <div className='page control-panel-login__form'>
        <Logo
          width='fit-content'
          viewBox='0 0 914 231'
          style={{
            height: '80px',
            paddingBottom: '40px',
          }}
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
