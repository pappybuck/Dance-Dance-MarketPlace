import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { Provider } from 'urql'
import { client } from '../lib/graphql'
import { Provider as Redux} from 'react-redux'
import { store } from '../Redux/store'
function MyApp({ Component, pageProps }: AppProps) {
  return (
    <Provider value={client}>
      <Redux store={store}>
        <Component {...pageProps} />
      </Redux>
    </Provider>
  )
}

export default MyApp
