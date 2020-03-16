import React from 'react'

import { render } from '@testing-library/react'

import Logo from './Logo'

describe('Logo', () => {
  it('should render', () => {
    const { container } = render(<Logo />)
    expect(container).toMatchSnapshot()
  })
})
