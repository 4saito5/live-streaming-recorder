// import fetch from 'node-fetch'
import axios from 'axios';
import { CheckList } from '../interfaces'


export async function getCheckListData(): Promise<CheckList[]> {
  try {
    const query = JSON.stringify({"query":"query findCheckList {\n checklists {\n id\n\t\t\tGroup\n\t\t\tName\n\t\t\tSite\n\t\t\tKey\n\t\t\tURL\n\t\t\tIsRecord\n OnLive\n }\n}\n"})

    const res = await axios(
    {
      url: 'http://gqlgen:5050/query',
      method: 'POST',
      headers: {
        'Accept-Encoding': 'gzip, deflate, br',
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Connection': 'keep-alive',
        'DNT': '1',
        'Origin': 'http://gqlgen:5050'
      },
      data: query
    })
    // // console.log(res)
    // console.log(res.status)
    // console.log(res.data.data)
    const items: CheckList[] = res.data.data.checklists
    return items
  } catch(e) {
    console.error(e)
  }
  return []
}

export async function createCheckListData(url: string) {
  if (!(/^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|www\.)[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$/.test(url))) {
    alert("URLを入力してください");
    return;
  }
  const res = await axios(
  {
    url: 'http://localhost:5050/add_url',
    method: 'POST',
    data: {
      'url': url
    }
  })
  location.reload(false)
  // console.log(res)
}

export async function deleteCheckListData(id: number) {
  try {
    const query = JSON.stringify({"query":"mutation DeleteCheckList {\n\tDeleteCheckList(input:{id: " + id + "}) {\n\t\tCode\n\t\tMessage\n\t}\n}"})

    const res = await axios(
    {
      url: 'http://localhost:5050/query',
      method: 'POST',
      headers: {
        // 'Accept-Encoding': 'gzip, deflate, br',
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        // 'Connection': 'keep-alive',
        // 'DNT': '1',
        // 'Origin': 'http://gqlgen:5050'
      },      
      data: query
    })
    location.reload(false)
    // console.log(res.data.data.DeleteCheckList.Code)
    // console.log(res.data.data.DeleteCheckList.Message)
    return true
  } catch(e) {
    console.error(e)
  }
  return false

  
}
