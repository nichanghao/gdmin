import request from '@/utils/axios'
import type {UnwrapNestedRefs} from "vue";

export type UserLoginVO = {
    username: string
    password: string
  }

// 登录
export const login = (data: UserLoginVO) => {
    return request.post({ url: '/login', data })
}

// 获取用户权限信息
export const getUserInfo = () => {
  return request.get({ url: '/sys/user/info'})
}