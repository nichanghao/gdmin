<template>
  <el-menu
    active-text-color="#ffd04b"
    background-color="#545c64"
    class="el-menu-vertical-demo"
    text-color="#fff"
    @open="handleOpen"
    @close="handleClose"
  >

    <template v-for="menu in routers">
        <el-sub-menu v-if="menu.children && menu.children.length > 0" :index=menu.path>
            <template #title>
                <el-icon><location /></el-icon>
                <span>{{ menu.name }}</span>
            </template>
            <el-menu-item-group>
                <el-menu-item v-for="item in menu.children" index="{{ item.path }}">{{ item.name }}</el-menu-item>
            </el-menu-item-group>
        </el-sub-menu>

        <el-menu-item v-else :index=menu.path>{{ menu.name }}</el-menu-item>
        
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

// 获取pinia中的菜单路由
const permissionStore = usePermissionStoreWithOut()
const {routers} = storeToRefs(permissionStore)
console.log(routers.value)

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
</script>
