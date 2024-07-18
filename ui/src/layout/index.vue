<template>
  <div class="common-layout">
    <el-container style="background-color:rgb(249 250 251 / 1);">
      <el-header style="padding: 0;"><Header /></el-header>
      <el-container>
        <el-aside width="200px"><Aside /></el-aside>
        <el-container>
          <el-main><Main /></el-main>
          <el-footer>Footer</el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import Main from './main/index.vue';
import Aside from './aside/index.vue';
import Header from './header/index.vue';
</script>
