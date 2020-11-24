import React, { ReactElement } from 'react'
import { AppProps } from 'next/app'

// import '../styles/app.scss'
import '../styles/index.css'

export default function App({ Component, pageProps }: AppProps): ReactElement {
  return <Component {...pageProps} />
}
