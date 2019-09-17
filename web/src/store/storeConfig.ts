import debounce from 'debounce-promise'
import {
    AnyAction, applyMiddleware, createStore as createReduxStore, Dispatch, Middleware,
    MiddlewareAPI,
} from 'redux'
import { composeWithDevTools as composeWithReduxDevTools } from 'redux-devtools-extension'
import reduxThunk from 'redux-thunk'

import { IhttpAPI } from '../httpAPI'
import reducer from '../reducers'

const logger = (store: MiddlewareAPI) => (next: Dispatch) => (
  action: AnyAction,
) => {
  console.log('action', action)
  return next(action)
}

export const setupStore = ({ httpAPI }: { httpAPI: IhttpAPI }) => {
  const middlewares: Middleware[] = []

  middlewares.push(
    reduxThunk.withExtraArgument({
      httpAPI: httpAPI,
      debounce,
    }),
  )

  if (process.env.NODE_ENV === 'development') {
    middlewares.push(logger)
  }

  return createReduxStore(
    reducer,
    composeWithReduxDevTools(applyMiddleware(...middlewares)),
  )
}
