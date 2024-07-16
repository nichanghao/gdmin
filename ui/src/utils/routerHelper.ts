import qs from 'qs'

const components = import.meta.glob('../views/**/*.vue')

// 路由生成
export const generateRoute = (routes: AppCustomRouteRecordRaw[]): AppRouteRecordRaw[] => {
    const res: AppRouteRecordRaw[] = []
    const modulesRoutesKeys = Object.keys(components)
    for (const route of routes) {
      // 1. 生成 meta 菜单元数据
      const meta = {
        title: route.name,
        icon: route.icon,
        hidden: !route.visible,
        noCache: !route.keepAlive,
        alwaysShow:
          route.children &&
          route.children.length === 1 &&
          (route.alwaysShow !== undefined ? route.alwaysShow : true)
      } as any
      // 特殊逻辑：如果后端配置的 MenuDO.component 包含 ?，则表示需要传递参数
      // 此时，我们需要解析参数，并且将参数放到 meta.query 中
      // 这样，后续在 Vue 文件中，可以通过 const { currentRoute } = useRouter() 中，通过 meta.query 获取到参数
      if (route.component && route.component.indexOf('?') > -1) {
        const query = route.component.split('?')[1]
        route.component = route.component.split('?')[0]
        meta.query = qs.parse(query)
      }
  
      // 2. 生成 data（AppRouteRecordRaw）
      let data: AppRouteRecordRaw = {
        path: route.path.indexOf('?') > -1 ? route.path.split('?')[0] : route.path,
        name:
          route.componentName && route.componentName.length > 0
            ? route.componentName
            : toCamelCase(route.path, true),
        redirect: route.redirect,
        meta: meta
      }
      // 顶级菜单路由
      if (!route.children && route.parentId == 0 && route.component) {
        // data.component = Layout
        data.meta = {}
        data.name = toCamelCase(route.path, true) + 'Parent'
        data.redirect = ''
        meta.alwaysShow = true
        const childrenData: AppRouteRecordRaw = {
          path: '',
          name:
            route.componentName && route.componentName.length > 0
              ? route.componentName
              : toCamelCase(route.path, true),
          redirect: route.redirect,
          meta: meta
        }
        const index = route?.component
          ? modulesRoutesKeys.findIndex((ev) => ev.includes(route.component))
          : modulesRoutesKeys.findIndex((ev) => ev.includes(route.path))
        childrenData.component = components[modulesRoutesKeys[index]]
        data.children = [childrenData]
      } else {
        // 非顶级菜单路由且有子节点路由
        if (route.children) {
          // data.component = Layout
          data.redirect = getRedirect(route.path, route.children)
        } else {
          // 叶子节点路由
          // 对后端传component组件路径和不传做兼容（如果后端传component组件路径，那么path可以随便写，如果不传，component组件路径会跟path保持一致）
          const index = route?.component
            ? modulesRoutesKeys.findIndex((ev) => ev.includes(route.component))
            : modulesRoutesKeys.findIndex((ev) => ev.includes(route.path))
          data.component = components[modulesRoutesKeys[index]]
        }
        if (route.children) {
          data.children = generateRoute(route.children)
        }
      }
      res.push(data as AppRouteRecordRaw)
    }
    return res
  }


// 处理路由重定向逻辑
export const getRedirect = (parentPath: string, children: AppCustomRouteRecordRaw[]):any => {
  if (!children || children.length == 0) {
    return parentPath
  }
  const path = generateRoutePath(parentPath, children[0].path)
  // 递归子节点
  if (children[0].children) 
    return getRedirect(path, children[0].children)
}
const generateRoutePath = (parentPath: string, path: string) => {
  if (parentPath.endsWith('/')) {
    parentPath = parentPath.slice(0, -1) // 移除默认的 /
  }
  if (!path.startsWith('/')) {
    path = '/' + path
  }
  return parentPath + path
}

// 中划线转驼峰
const toCamelCase = (str: string, upperCaseFirst: boolean) => {
  str = (str || '')
    .replace(/-(.)/g, function (group1: string) {
      return group1.toUpperCase()
    })
    .replaceAll('-', '')

  if (upperCaseFirst && str) {
    str = str.charAt(0).toUpperCase() + str.slice(1)
  }

  return str
}