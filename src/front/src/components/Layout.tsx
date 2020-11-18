import React, { ReactNode } from 'react'
import Head from 'next/head'

import { createCheckListData } from '../lib/checklists'


type Props = {
  children?: ReactNode
  title?: string
  url?: string
}

// import { useState } from "react";
// // const [todos, setTodos] = useState([]);
// const [tmpTodo, setTmpTodo] = useState("")
// const setUrl = (value: string) => {
//   url = value
// }
const addUrl = (url: string) => {
  // console.log(url)
  createCheckListData(url)
};

const Layout = ({ children, title = 'This is the default title', url }: Props) => (
  <div>
    <Head>
      <title>{title}</title>
      <meta charSet="utf-8" />
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
    </Head>
    <header>

      <div className="form">
        <input className="url"
          type="text"
          name="url"
          onChange={e => {url = e.target.value}}
          value={url}
          style={{width:'80%'}}
        />
        <button onClick={() => addUrl(url)}>Add</button>
      </div>


    </header>
    {children}
    <footer>
      <hr />
      <span>I'm here to stay (Footer)</span>
    </footer>
  </div>
)

export default Layout
