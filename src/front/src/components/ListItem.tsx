import React from 'react'

import { CheckList } from '../interfaces'

type Props = {
  data: CheckList
}

const ListItem = ({ data }: Props) => (
  <span className="">
    <span className="cl-group">{data.Group}</span>
    <span className="cl-name">{data.Name}</span>
    <span className="cl-site">{data.Site}</span>
    <span className="cl-url">{data.URL}</span>
    <span className="cl-is-record">{data.IsRecord}</span>
    <span className="cl-on-live">{data.OnLive}</span>
    <span className="cl-edit"><a href="#" className="text-indigo-600 hover:text-indigo-900">Edit</a></span>
  </span>
)

export default ListItem
