import React from 'react'

import { render } from '@testing-library/react'

import Sidebar from './Sidebar'

describe('Sidebar', () => {
  it('should render', () => {
    const { container } = render(<Sidebar />)
    expect(container).toMatchSnapshot()
  })
})
