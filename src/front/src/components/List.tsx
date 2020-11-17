import * as React from 'react'
import ListItem from './ListItem'
import { CheckList } from '../interfaces'

type Props = {
  items: CheckList[]
}

const List = ({ items }: Props) => (
  <ul>
    {items.map((item) => (
      <li key={item.id}>
        <ListItem data={item} />
      </li>
    ))}
  </ul>
)

export default List
