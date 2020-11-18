// import fetch from 'node-fetch'
import axios from 'axios';
import { CheckList } from '../interfaces'


export async function getCheckListData(): Promise<CheckList[]> {
  try {
    const query = JSON.stringify({"query":"query findCheckList {\n checklists {\n id\n\t\t\tGroup\n\t\t\tName\n\t\t\tSite\n\t\t\tKey\n\t\t\tURL\n\t\t\tIsRecord\n }\n}\n"})

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
    // // console.log(res);
    // console.log(res.status);
    // console.log(res.data.data);
    const items: CheckList[] = res.data.data.checklists
    return items
  } catch(e) {
    console.log(e);
  }
  return []
}

export async function createCheckListData(url: string) {
  console.log(url)
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

  location.reload(false);


  // const res = await axios.post('http://localhost:5050/add_url',
  //   {
  //     'url': url
  //   },
  //   {
  //     headers: {
  //       'Access-Control-Allow-Origin': '*',
  //     }
  //   }
  // )


  console.log(res)



}
// export async function createCheckListData(): Promise<CheckList[]> {
//   try {
//     const query = JSON.stringify({"query":"mutation CreateCheckList {\n CreateCheckList(input:{Group: \"team8\", Name: \"大西桃香\", Site: \"Momoka_Oonishi\", URL: \"https://www.showroom-live.com/\", IsRecord: 1}) {\n id\n\t\tGroup\n\t\tName\n\t\tSite\n\t\tKey\n\t\tIsRecord\n }\n}\n"})

//     const res = await axios(
//     {
//       url: 'http://gqlgen:5050/query',
//       method: 'POST',
//       headers: {
//         'Accept-Encoding': 'gzip, deflate, br',
//         'Content-Type': 'application/json',
//         'Accept': 'application/json',
//         'Connection': 'keep-alive',
//         'DNT': '1',
//         'Origin': 'http://gqlgen:5050'
//       },
//       data: query
//     })
//     // console.log(res);
//     console.log(res.status);
//     console.log(res.data.data);
//     const items: CheckList[] = res.data.data.CreateCheckList
//     return items
//   } catch(e) {
//     console.log(e);
//   }
//   return []
// }
