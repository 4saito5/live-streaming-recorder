import React from 'react'

import { CheckList } from '../interfaces'

type Props = {
  data: CheckList
}

const ListItem = ({ data }: Props) => (
    <a>
      {data.id}: {data.Group}, {data.Name}, {data.Site}, {data.Key}, {data.URL}, {data.IsRecord}, {data.OnLive}
    </a>
)

export default ListItem
