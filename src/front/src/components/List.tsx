import * as React from 'react'
import ListItem from './ListItem'
import { CheckList } from '../interfaces'

type Props = {
  items: CheckList[]
}

const List = ({ items }: Props) => (
  <ul className="items-center p-8 list-none ">
    <li>
      <span className="bg-gray-100">
        <span className="cl-group">Group</span>
        <span className="cl-name">Name</span>
        <span className="cl-site">Site</span>
        <span className="cl-url">URL</span>
        <span className="cl-is-record">IsRec</span>
        <span className="cl-on-live">OnLive</span>
      </span>
    </li>

    {items.map((item) => (
      <li key={item.id}>
        <ListItem data={item} />
      </li>
    ))}
  </ul>
)

export default List
