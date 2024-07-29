import {request} from '../request';

/** get role list */
export function fetchGetRoleList(params?: Api.SystemManage.RoleSearchParams) {
  return request<Api.SystemManage.RoleList>({
    url: '/sys/role/page',
    method: 'post',
    data: params
  });
}

/** add role */
export function addRole(data: any) {
  return request<any>({
    url: '/sys/role/add',
    method: 'post',
    data: data
  });
}

/** edit role */
export function editRole(data: any) {
  return request<any>({
    url: '/sys/role/edit',
    method: 'put',
    data: data
  });
}

/** delete role */
export function deleteRole(id: number) {
  return request<any>({
    url: '/sys/role/delete',
    method: 'delete',
    params: {
      id
    }
  });
}


/**
 * get all roles
 *
 * these roles are all enabled
 */
export function fetchGetAllRoles() {
  return request<Api.SystemManage.AllRole[]>({
    url: '/systemManage/getAllRoles',
    method: 'get'
  });
}

/** get user list */
export function fetchGetUserList(params?: Api.SystemManage.UserSearchParams) {
  return request<Api.SystemManage.UserList>({
    url: '/systemManage/getUserList',
    method: 'get',
    params
  });
}

/** get menu list */
export function fetchGetMenuList() {
  return request<Api.SystemManage.MenuList>({
    url: '/sys/menu/tree',
    method: 'get'
  });
}

/** get all pages */
export function fetchGetAllPages() {
  return request<string[]>({
    url: '/systemManage/getAllPages',
    method: 'get'
  });
}

/** get menu tree */
export function fetchGetMenuTree() {
  return request<Api.SystemManage.MenuTree[]>({
    url: '/sys/menu/all-simple-tree',
    method: 'get'
  });
}

/** get menu by role */
export function fetchGetMenuByRole(id: number) {
  return request<number[]>({
    url: '/sys/menu/list-by-role',
    method: 'get',
    params: {
      id
    }
  });
}

/** assign menu to role */
export function assignMenuToRole(menuIds: number[], roleId: number) {
  return request<any>({
    url: '/sys/role/assign-menus',
    method: 'put',
    data: {
      menuIds,
      roleId
    }
  });
}


/** add menu data */
export function addMenu(menu: any) {
  return request<any>({
    url: '/sys/menu/add',
    method: 'post',
    data: menu
  });
}

/** edit menu data */
export function editMenu(menu: any) {
  return request<any>({
    url: '/sys/menu/edit',
    method: 'put',
    data: menu
  });
}

/** delete menu data */
export function deleteMenu(id: number) {
  return request<any>({
    url: '/sys/menu/delete',
    method: 'delete',
    params: {
      id
    }
  });
}
