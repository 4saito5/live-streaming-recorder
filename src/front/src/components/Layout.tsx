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
  createCheckListData(url)
};

const keyPress = (e: any) => {
  if (e.which === 13) {
    addUrl(e.target.value)
  }
}

const Layout = ({ children, title = 'This is the default title', url }: Props) => (
  <div>
    <Head>
      <title>{title}</title>
      <meta charSet="utf-8" />
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
    </Head>
    <header>

      <div className="form items-center p-8">
        <input className="border-2 px-2 py-1 rounded "
          type="text"
          name="url"
          onChange={e => {url = e.target.value}}
          value={url}
          placeholder="https://www.showroom-live.com/xxxxx"
          onKeyPress={(e) => keyPress(e)}
          style={{width:'80%'}}
        />
        <button className="border-2 px-2 py-1  rounded" onClick={() => addUrl(url!)}>Add</button>
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
