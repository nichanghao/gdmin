import { useCache } from '@/hooks/useCache';

const { wsCache } = useCache();

const AccessTokenKey = 'ACCESS_TOKEN';
const LoginFormKey = 'LOGIN_FORM';

// 获取token
export const getAccessToken = () => {
  return wsCache.get(AccessTokenKey) ? wsCache.get(AccessTokenKey) : '';
};

// 设置token
export const setToken = (token: string) => {
  wsCache.set(AccessTokenKey, token);
};

// 删除token
export const removeToken = () => {
  wsCache.delete(AccessTokenKey);
};

/** 格式化token（jwt格式） */
export const formatToken = (token: string): string => {
  return 'Bearer ' + token;
};
// ========== 账号相关 ==========
export type LoginFormType = {
  username: string;
  password: string;
  rememberMe: boolean;
};

export const getLoginForm = () => {
  const loginForm: LoginFormType = wsCache.get(LoginFormKey);
  return loginForm;
};

export const setLoginForm = (loginForm: LoginFormType) => {
  wsCache.set(LoginFormKey, loginForm, { exp: 7 * 24 * 60 * 60 });
};

export const removeLoginForm = () => {
  wsCache.delete(LoginFormKey);
};
