import { store } from '@/stores';
import { defineStore } from 'pinia';
import { getAccessToken, removeToken } from '@/utils/auth';
import { CACHE_KEY, useCache, deleteUserCache } from '@/hooks/useCache';
import { getUserInfo } from '@/api/user';

const { wsCache } = useCache();

interface UserVO {
  id: number;
  avatar: string;
  nickname: string;
}

interface UserInfoVO {
  // USER 缓存
  permissions: string[];
  roles: string[];
  isSetUser: boolean;
  user: UserVO;
}

export const useUserStore = defineStore('admin-user', {
  state: (): UserInfoVO => ({
    permissions: [],
    roles: [],
    isSetUser: false,
    user: {
      id: 0,
      avatar: '',
      nickname: ''
    }
  }),
  getters: {
    getPermissions(): string[] {
      return this.permissions;
    },
    getRoles(): string[] {
      return this.roles;
    },
    getIsSetUser(): boolean {
      return this.isSetUser;
    },
    getUser(): UserVO {
      return this.user;
    }
  },
  actions: {
    async setUserInfoAction() {
      if (!getAccessToken()) {
        this.resetState();
        return null;
      }
      let userInfo = wsCache.get(CACHE_KEY.USER);
      if (!userInfo) {
        userInfo = await getUserInfo();
      }
      this.permissions = userInfo.permissions;
      this.user = userInfo.user;
      this.roles = userInfo.user.roles;
      this.isSetUser = true;
      wsCache.set(CACHE_KEY.USER, userInfo);
      wsCache.set(CACHE_KEY.ROLE_ROUTERS, userInfo.menuTree);
    },
    async setUserAvatarAction(avatar: string) {
      const userInfo = wsCache.get(CACHE_KEY.USER);
      // NOTE: 是否需要像`setUserInfoAction`一样判断`userInfo != null`
      this.user.avatar = avatar;
      userInfo.user.avatar = avatar;
      wsCache.set(CACHE_KEY.USER, userInfo);
    },
    async setUserNicknameAction(nickname: string) {
      const userInfo = wsCache.get(CACHE_KEY.USER);
      // NOTE: 是否需要像`setUserInfoAction`一样判断`userInfo != null`
      this.user.nickname = nickname;
      userInfo.user.nickname = nickname;
      wsCache.set(CACHE_KEY.USER, userInfo);
    },
    async loginOut() {
      removeToken();
      deleteUserCache(); // 删除用户缓存
      this.resetState();
    },
    resetState() {
      this.permissions = [];
      this.roles = [];
      this.isSetUser = false;
      this.user = {
        id: 0,
        avatar: '',
        nickname: ''
      };
    }
  }
});

export const useUserStoreWithOut = () => {
  return useUserStore(store);
};
