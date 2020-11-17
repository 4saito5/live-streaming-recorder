// import fetch from 'node-fetch'
import axios from 'axios';
import { CheckList } from '../interfaces'


export async function getCheckListData(): Promise<CheckList[]> {
  try {
    const query = JSON.stringify({"query":"query findCheckList {\n checklists {\n id\n\t\t\tGroup\n\t\t\tName\n\t\t\tSite\n\t\t\tURL\n\t\t\tIsRecord\n }\n}"})
    // const query = JSON.stringify({"query":"mutation CreateCheckList {\n CreateCheckList(input:{Group: \"team8\", Name: \"大西桃香\", Site: \"Momoka_Oonishi\", URL: \"https://www.showroom-live.com/\", IsRecord: 1}) {\n id\n\t\tGroup\n\t\tName\n\t\tSite\n\t\tURL\n\t\tIsRecord\n }\n}\n"})

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
    // console.log(res);
    console.log(res.status);
    console.log(res.data.data);
    const items: CheckList[] = res.data.data.checklists
    return items
  } catch(e) {
    console.log(e);
  }
  return []
}

export async function createCheckListData(): Promise<CheckList[]> {
  try {
    const query = JSON.stringify({"query":"mutation CreateCheckList {\n CreateCheckList(input:{Group: \"team8\", Name: \"大西桃香\", Site: \"Momoka_Oonishi\", URL: \"https://www.showroom-live.com/\", IsRecord: 1}) {\n id\n\t\tGroup\n\t\tName\n\t\tSite\n\t\tURL\n\t\tIsRecord\n }\n}\n"})

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
    // console.log(res);
    console.log(res.status);
    console.log(res.data.data);
    const items: CheckList[] = res.data.data.CreateCheckList
    return items
  } catch(e) {
    console.log(e);
  }
  return []
}
