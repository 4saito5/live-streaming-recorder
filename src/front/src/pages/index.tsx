// import fetch from 'node-fetch'
// import axios from 'axios';
import { GetStaticProps } from 'next'

import { CheckList } from '../interfaces'
// import { sampleCheckListData } from '../utils/sample-check-data'
import { getCheckListData } from '../lib/checklists'
import Layout from '../components/Layout'
import List from '../components/List'


type Props = {
  items: CheckList[]
  // url: string
}

const IndexPage = ({ items }: Props) => (

  <Layout title="Home | Next.js + TypeScript Example">
    <List items={items} />
  </Layout>
)

export const getStaticProps: GetStaticProps = async () => {
  // Example for including static props in a Next.js function component page.
  // Don't forget to include the respective types for any props passed into
  // the component.

  const items: CheckList[] = await getCheckListData()
  return { props: { items } }
}

export default IndexPage
