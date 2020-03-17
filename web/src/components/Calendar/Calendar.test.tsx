import React from 'react'

import { fireEvent, render } from '@testing-library/react'

import Calendar from './Calendar'


describe('Calendar', () => {
  it('should render', () => {
    const { container } = render(
      <Calendar>
      </Calendar>)
    expect(container).toMatchSnapshot()
  })
})
