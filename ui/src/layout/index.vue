<template>
    <div class="common-layout">
      <el-container>
        <el-aside width="200px">Aside</el-aside>
        <el-container>
          <el-header>Header</el-header>
          <el-main>
            <main/>
          </el-main>
          <el-footer>Footer</el-footer>
        </el-container>
      </el-container>
    </div>
  </template>

<script setup>
    import main from './main/index.vue'
</script>