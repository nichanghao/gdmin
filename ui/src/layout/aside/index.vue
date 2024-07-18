<template>
  <el-menu
    active-text-color="#ffd04b"
    background-color="#545c64"
    class="el-menu-vertical-demo"
    text-color="#fff"
    @open="handleOpen"
    @close="handleClose"
  >
    <el-menu-item @click="handleRouter({path: '/home'})" index="/home">
      <template #title>
        <el-icon><i-ep-house /></el-icon>
        <span>首页</span>
      </template>
    </el-menu-item>

    <template v-for="menu in menuTree">
        <el-sub-menu v-if="menu.children && menu.children.length > 0" :index=menu.path>
            <template #title>
                <el-icon><location /></el-icon>
                <span>{{ menu.meta.title }}</span>
            </template>
            <el-menu-item-group>
                <el-menu-item v-for="item in menu.children" @click="handleRouter(item)" index="{{ item.path }}">{{ item.meta.title }}</el-menu-item>
            </el-menu-item-group>
        </el-sub-menu>

        <el-menu-item v-else @click="handleRouter(item)" :index=menu.path>{{ menu.meta.title }}</el-menu-item>
        
    </template>


    <el-sub-menu index="1">
      <template #title>
        <el-icon><location /></el-icon>
        <span>Navigator One</span>
      </template>
      <el-menu-item-group title="Group One">
        <el-menu-item index="1-1">item one</el-menu-item>
        <el-menu-item index="1-2">item two</el-menu-item>
      </el-menu-item-group>
      <el-menu-item-group title="Group Two">
        <el-menu-item index="1-3">item three</el-menu-item>
      </el-menu-item-group>
      <el-sub-menu index="1-4">
        <template #title>item four</template>
        <el-menu-item index="1-4-1">item one</el-menu-item>
      </el-sub-menu>
    </el-sub-menu>
    <el-menu-item index="2">
      <el-icon><icon-menu /></el-icon>
      <span>Navigator Two</span>
    </el-menu-item>
    <el-menu-item index="3" disabled>
      <el-icon><document /></el-icon>
      <span>Navigator Three</span>
    </el-menu-item>
    <el-menu-item index="4">
      <el-icon><setting /></el-icon>
      <span>Navigator Four</span>
    </el-menu-item>
  </el-menu> 
</template>

<script lang="ts" setup>
import { Document, Menu as IconMenu, Location, Setting } from '@element-plus/icons-vue';
import { usePermissionStoreWithOut } from '@/stores/permission';
import { storeToRefs } from 'pinia';
import router from '@/router';


// 获取pinia中的菜单路由
const permissionStore = usePermissionStoreWithOut()
const {menuTree} = storeToRefs(permissionStore)
console.log("menuTree: ",menuTree.value)


// 切换路由
function handleRouter(menu) {
console.log("handleRouter: ",menu)

  router.push(menu.path);
}



const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
</script>
