import { defineStore } from 'pinia';
import { store } from '@/stores';
import { CACHE_KEY, useCache } from '@/hooks/useCache';
import { flatMultiLevelRoutes, generateRoute } from '@/utils/routerHelper';
import { cloneDeep } from 'lodash-es';
import { routerArray } from '@/router/index';

const { wsCache } = useCache();

export interface PermissionState {
  routers: AppRouteRecordRaw[];
  addRouters: AppRouteRecordRaw[];
  menuTabRouters: AppRouteRecordRaw[];
  menuTree: AppRouteRecordRaw[];
}

export const usePermissionStore = defineStore('permission', {
  state: (): PermissionState => ({
    routers: [],
    addRouters: [],
    menuTabRouters: [],
    // 左侧导航菜单树
    menuTree: []
  }),
  getters: {
    getRouters(): AppRouteRecordRaw[] {
      return this.routers;
    },
    getAddRouters(): AppRouteRecordRaw[] {
      return flatMultiLevelRoutes(cloneDeep(this.addRouters))
    },
    getMenuTabRouters(): AppRouteRecordRaw[] {
      return this.menuTabRouters;
    },
    getMenuTree(): AppRouteRecordRaw[] {
      return this.menuTree;
    }
  },
  actions: {
    async generateRoutes(): Promise<unknown> {
      return new Promise<void>(async (resolve) => {
        let res: AppCustomRouteRecordRaw[] = [];
        if (wsCache.get(CACHE_KEY.ROLE_ROUTERS)) {
          res = wsCache.get(CACHE_KEY.ROLE_ROUTERS) as AppCustomRouteRecordRaw[];
        }
        const routerMap: AppRouteRecordRaw[] = generateRoute(res);
        this.menuTree = routerMap;

        // 动态路由，404一定要放到最后面
        this.addRouters = routerMap.concat([
          {
            path: '/:path(.*)*',
            redirect: '/404',
            name: '404Page',
            meta: {
              hidden: true,
              breadcrumb: false
            }
          }
        ]);
        // 渲染菜单的所有路由
        this.routers = cloneDeep(routerArray).concat(routerMap);
        resolve();
      });
    },
    setMenuTabRouters(routers: AppRouteRecordRaw[]): void {
      this.menuTabRouters = routers;
    }
  },
  persist: false
});

export const usePermissionStoreWithOut = () => {
  return usePermissionStore(store);
};
