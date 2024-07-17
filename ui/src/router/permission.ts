import router from '@/router';
import type { RouteRecordRaw } from 'vue-router';
import { getAccessToken } from '@/utils/auth';
import { useUserStoreWithOut } from '@/stores/user';
import { usePermissionStoreWithOut } from '@/stores/permission';

// 路由不重定向白名单
const whiteList = ['/login'];

// 路由加载前
router.beforeEach(async (to, from, next) => {
  if (getAccessToken()) {
    if (to.path === '/login') {
      next({ path: '/' });
    } else {
      const userStore = useUserStoreWithOut();
      const permissionStore = usePermissionStoreWithOut();

      if (!userStore.getIsSetUser) {
        // 设置用户信息
        await userStore.setUserInfoAction();

        // 生成路由
        await permissionStore.generateRoutes();
        permissionStore.getAddRouters.forEach((route) => {
          router.addRoute(route as unknown as RouteRecordRaw) // 动态添加可访问路由表
        })
        const redirectPath = from.query.redirect || to.path;
        // 跳转时带参数
        const redirect = decodeURIComponent(redirectPath as string);
        const { paramsObject: query } = parseURL(redirect);
        const nextData =
          to.path === redirect ? { ...to, replace: true } : { path: redirect, query };
        next(nextData);
      } else {
        next();
      }
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next();
    } else {
      next(`/login?redirect=${to.fullPath}`); // 否则全部重定向到登录页
    }
  }
});

router.afterEach((to) => {});

const parseURL = (
  url: string | null | undefined
): { basePath: string; paramsObject: { [key: string]: string } } => {
  // 如果输入为 null 或 undefined，返回空字符串和空对象
  if (url == null) {
    return { basePath: '', paramsObject: {} };
  }

  // 找到问号 (?) 的位置，它之前是基础路径，之后是查询参数
  const questionMarkIndex = url.indexOf('?');
  let basePath = url;
  const paramsObject: { [key: string]: string } = {};

  // 如果找到了问号，说明有查询参数
  if (questionMarkIndex !== -1) {
    // 获取 basePath
    basePath = url.substring(0, questionMarkIndex);

    // 从 URL 中获取查询字符串部分
    const queryString = url.substring(questionMarkIndex + 1);

    // 使用 URLSearchParams 遍历参数
    const searchParams = new URLSearchParams(queryString);
    searchParams.forEach((value, key) => {
      // 封装进 paramsObject 对象
      paramsObject[key] = value;
    });
  }

  // 返回 basePath 和 paramsObject
  return { basePath, paramsObject };
};
