import React from 'react'

import { fireEvent, render } from '@testing-library/react'

import Button from './Button'

describe('Button', () => {
  it('should render', () => {
    const { container } = render(<Button />)
    expect(container).toMatchSnapshot()
  })
})
