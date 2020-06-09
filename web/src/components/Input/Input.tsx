import React, { InputHTMLAttributes, useEffect, useRef, useState } from 'react'

interface IProps extends InputHTMLAttributes<HTMLInputElement> {
  required?: boolean
  secret?: boolean
  number?: boolean
}
const Input: React.FC<IProps> = props => {
  const {
    onChange,
    onKeyPress,
    placeholder,
    required,
    secret,
    number,
    className,
    id,
  } = props

  const [ isActive, setIsActive ] = useState<boolean>(false)
  const [ text, setText ] = useState<string>('')

  const inputRef = useRef<HTMLInputElement>(null)

  useEffect(
    () => {
      if (isActive && inputRef.current) {
        inputRef.current.focus()
      } else if (inputRef.current) {
        inputRef.current.blur()
      }
    },
    [ isActive ],
  )

  const onPlaceholderClick = (): void => {
    if (!isActive) {
      setIsActive(true)
    }
  }

  const onPlaceholderBlur = (): void => {
    const isTextEmpty = text === ''
    if (isActive && isTextEmpty) {
      setIsActive(false)
    }
  }

  const onInputChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    const eventText = e.currentTarget.value
    setText(eventText)
    const isTextEmpty = eventText === ''
    if (isTextEmpty && !isActive) {
      setIsActive(false)
    } else {
      setIsActive(true)
    }

    if (onChange) {
      onChange(e)
    }
  }

  const onInputFocus = (): void => {
    if (!isActive) {
      setIsActive(true)
    }
  }

  const onInputBlur = (e: React.FocusEvent<HTMLInputElement>): void => {
    const eventText = e.currentTarget.value
    setText(eventText)
    const isTextEmpty = eventText === ''
    if (isActive && isTextEmpty) {
      setIsActive(false)
    }
  }

  let inputType = 'text'
  if (number) {
    inputType = 'number'
  }
  if (secret) {
    inputType = 'password'
  }

  return (
    <div
      className={`
        input
        ${isActive ? 'input_active' : ''}
        ${className}
      `}
      id={id}
    >
      <input
        className={`
          input__element
        `}
        ref={inputRef}
        required={required}
        type={inputType}
        onChange={onInputChange}
        onKeyPress={onKeyPress}
        onFocus={onInputFocus}
        onBlur={onInputBlur}
      />
      <div
        className='input__placeholder'
        onClick={onPlaceholderClick}
        onBlur={onPlaceholderBlur}
      >
        {placeholder}
      </div>
    </div>
  )
}

export default Input
