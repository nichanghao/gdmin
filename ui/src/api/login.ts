import request from '@/utils/axios'

export type UserLoginVO = {
    username: string
    password: string
  }

// 登录
export const login = (data: UserLoginVO) => {
    return request.post({ url: '/login', data })
  }