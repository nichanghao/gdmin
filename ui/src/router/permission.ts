import router from '@/router'
import type { RouteRecordRaw } from 'vue-router'
import { getAccessToken } from '@/utils/auth'
import { useUserStoreWithOut } from '@/stores/user'
import { usePermissionStoreWithOut } from '@/stores/permission'



// 路由不重定向白名单
const whiteList = [
  '/login',
]

// 路由加载前
router.beforeEach(async (to, from, next) => {
  if (getAccessToken()) {
    if (to.path === '/login') {
      next({ path: '/' })
    } else {
        const userStore = useUserStoreWithOut()
        const permissionStore = usePermissionStoreWithOut()


      if (!userStore.getIsSetUser) {

        await userStore.setUserInfoAction()

        // 后端过滤菜单
        await permissionStore.generateRoutes()
        // permissionStore.getAddRouters.forEach((route) => {
        //   router.addRoute(route as unknown as RouteRecordRaw) // 动态添加可访问路由表
        // })
        const redirectPath = from.query.redirect || to.path
        // 修复跳转时不带参数的问题
        const redirect = decodeURIComponent(redirectPath as string)
        // const { basePath, paramsObject: query } = parseURL(redirect)
        // const nextData = to.path === redirect ? { ...to, replace: true } : { path: redirect, query }
        // next(nextData)
      } else {
        next()
      }
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      next(`/login?redirect=${to.fullPath}`) // 否则全部重定向到登录页
    }
  }
})

router.afterEach((to) => {
  
})
