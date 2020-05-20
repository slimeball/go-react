import env from './env'
import axios from 'axios'

export default async (type, url, data, headers) => {
  type = type.toLowerCase() || 'get';
  url = env.baseUrl + env.version + url;
  if (type === 'get') {
    return new Promise((resolve, reject) => {
      let params = { params: data }
      if (headers) {
        params.headers = headers
      }
      axios.get(url, params).then(response => {
        resolve(response.data)
      }).then(response => {
        reject(response)
      })
    })
  } else if(type === 'post'){
    return new Promise((resolve, reject) => {
      axios.post(url, data).then(response => {
        resolve(response.data);
      }, response => {
        reject(response);
      });
    })
  }
}
