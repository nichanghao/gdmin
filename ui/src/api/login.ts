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