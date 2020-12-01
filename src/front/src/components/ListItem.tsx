import React from 'react'

import { CheckList } from '../interfaces'
import { deleteCheckListData } from '../lib/checklists'

type Props = {
  data: CheckList
}


const deleteTask = (id: number) => {
  // console.log(url)
  deleteCheckListData(id)
};


const ListItem = ({ data }: Props) => (
  <span className="">
    <span className="cl-id">{data.id}</span>
    <span className="cl-group">{data.Group}</span>
    <span className="cl-name">{data.Name}</span>
    <span className="cl-site">{data.Site}</span>
    <span className="cl-url">{data.URL}</span>
    <span className="cl-is-record">{data.IsRecord}</span>
    <span className="cl-on-live">{data.OnLive}</span>
    {/* <span className="cl-edit"><a href="#" className="text-indigo-600 hover:text-indigo-900">Edit</a></span> */}
    <span className="cl-del"><a href="#" className="text-indigo-600 hover:text-indigo-900" onClick={() => deleteTask(data.id!)}>Del</a></span>
  </span>
)

export default ListItem
