import axios, { type AxiosInstance, type InternalAxiosRequestConfig } from 'axios';

const request: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL,
  // 超时时间10s，单位是毫秒
  timeout: 10000,
  // 跨域请求时是否携带cookie
  withCredentials: false
});

// 不需要token认证的白名单
let whiteSet:Set<string> = new Set(['/login'])

// 请求拦截器
request.interceptors.request.use((config) => {

  let isToken = true
  if (config.url && whiteSet.has(config.url)) {
    isToken = false
  }

  if (isToken) {
    // todo
    config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`
  }

  return config
}, (error) => {
  console.log(error)
  Promise.reject(error)
})


// 响应拦截器
request.interceptors.response.use((res) => {
  if (!res) {
    // 未响应时
    throw new Error('Network Error')
  }

  return res
}, (error) => {

})