import React from 'react'

import { render } from '@testing-library/react'

import Navigation from './Navigation'

describe('Navigation', () => {
  it('should render', () => {
    const { container } = render(
      <Navigation>
        <Navigation.Button href='#' label='hi' />
      </Navigation>
    )
    expect(container).toMatchSnapshot()
  })
})
