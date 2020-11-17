import React from 'react'

import { CheckList } from '../interfaces'

type Props = {
  data: CheckList
}

const ListItem = ({ data }: Props) => (
    <a>
      {data.id}: {data.Group}, {data.Name}, {data.Site}, {data.URL}, {data.IsRecord}
    </a>
)

export default ListItem
