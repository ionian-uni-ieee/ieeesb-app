import React from 'react'

import { fireEvent, render } from '@testing-library/react'

import Input from './Input'

describe('Input', () => {
  it('should render', () => {
    const { container } = render(<Input />)
    expect(container).toMatchSnapshot()
  })
  it('should get an input', () => {
    const { container } = render(<Input />)
    const input = container.querySelector('input') as HTMLInputElement
    expect(input).not.toBe(null)
  })
  it('should write input', () => {
    const { container } = render(<Input />)
    const input = container.querySelector('input') as HTMLInputElement
    fireEvent.input(input, { target: { value: 'test' } })
    expect(input.value).toBe('test')
  })
})
