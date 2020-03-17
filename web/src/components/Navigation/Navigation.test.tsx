import React from 'react'
import { MemoryRouter } from 'react-router'

import { render } from '@testing-library/react'

import Navigation from './Navigation'

describe('Navigation', () => {
  it('should render', () => {
    const { container } = render(
      <MemoryRouter>
        <Navigation>
          <Navigation.Button href='#' label='hi' />
        </Navigation>
      </MemoryRouter>
    )
    expect(container).toMatchSnapshot()
  })
})
