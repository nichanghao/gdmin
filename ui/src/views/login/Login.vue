<template>
  <div class="login-container">
    <el-form
      label-position="left"
      label-width="auto"
      :model="loginFormData"
      style="max-width: 600px; margin: 40vh auto"
    >
      <el-form-item>
        <el-input v-model="loginFormData.name" placeholder="请输入用户名">
          <template #prefix>
            <i-ep-user />
          </template>
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-input
          v-model="loginFormData.region"
          type="password"
          placeholder="请输入密码"
          show-password
        >
          <template #prefix>
            <i-ep-lock />
          </template>
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-row justify="space-between" style="width: 100%">
          <el-col :span="8">
            <el-checkbox v-model="loginFormData.rememberMe" label="记住我"></el-checkbox>
          </el-col>
          <el-col :offset="6" :span="8">
            <el-link style="float: right" type="primary">忘记密码？</el-link>
          </el-col>
        </el-row>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';

const loginFormData = reactive({
  name: '',
  region: '',
  rememberMe: ''
});
</script>

<style scoped></style>
