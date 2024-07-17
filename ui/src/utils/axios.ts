import axios, { type AxiosInstance, type AxiosResponse } from 'axios';
import { getAccessToken } from '@/utils/auth';

const axiosInstance: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL,
  // 超时时间10s，单位是毫秒
  timeout: 10000,
  // 跨域请求时是否携带cookie
  withCredentials: false
});

// 不需要token认证的白名单
let whiteSet: Set<string> = new Set(['/login']);

// 请求拦截器
axiosInstance.interceptors.request.use(
  (config) => {
    let isToken = true;
    if (config.url && whiteSet.has(config.url)) {
      isToken = false;
    }

    if (isToken) {
      config.headers.Authorization = `Bearer ` + getAccessToken();
    }

    return config;
  },
  (error) => {
    console.log(error);
    Promise.reject(error);
  }
);

// 响应拦截器
axiosInstance.interceptors.response.use(
  (res: AxiosResponse<any>) => {
    if (!res) {
      // 未响应时
      throw new Error('Network Error');
    }

    const { code, msg } = res.data;
    if (code == 201) {
      // 提示消息
      ElMessage.error(msg);
    } else if (code == 401) {
      // token失效
      // 已经在登录页面
      if (window.location.href.includes('login?redirect=')) {
        return Promise.reject('token失效');
      }

      ElMessageBox.confirm('登录信息已过期，请重新登录', '提示', {
        showCancelButton: false,
        closeOnClickModal: false,
        showClose: false,
        confirmButtonText: '确定',
        type: 'warning'
      }).then(() => {
        window.location.href = window.location.href;
      });
    }

    return res.data;
  },
  (error) => {
    console.log('response error: ', error);

    ElMessage.error('网络错误，请稍后重试');

    return Promise.reject('Network Error');
  }
);

const request = (option: any) => {
  const { url, method, params, data, headersType, responseType, ...config } = option;
  return axiosInstance({
    url: url,
    method,
    params,
    data,
    ...config,
    responseType: responseType,
    headers: {
      'Content-Type': headersType || 'application/json'
    }
  });
};

export default {
  get: async <T = any>(option: any) => {
    const res = await request({ method: 'GET', ...option });
    return res.data as unknown as T;
  },
  post: async <T = any>(option: any) => {
    const res = await request({ method: 'POST', ...option });
    return res.data as unknown as T;
  },
  postOriginal: async (option: any) => {
    const res = await request({ method: 'POST', ...option });
    return res;
  },
  delete: async <T = any>(option: any) => {
    const res = await request({ method: 'DELETE', ...option });
    return res.data as unknown as T;
  },
  put: async <T = any>(option: any) => {
    const res = await request({ method: 'PUT', ...option });
    return res.data as unknown as T;
  },
  download: async <T = any>(option: any) => {
    const res = await request({ method: 'GET', responseType: 'blob', ...option });
    return res as unknown as Promise<T>;
  },
  upload: async <T = any>(option: any) => {
    option.headersType = 'multipart/form-data';
    const res = await request({ method: 'POST', ...option });
    return res as unknown as Promise<T>;
  }
};
