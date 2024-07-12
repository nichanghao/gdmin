<template>
  <div class="login-container">
    <el-form
      label-position="left"
      label-width="auto"
      :model="loginForm.FormData"
      style="max-width: 600px; margin: 40vh auto"
    >
      <el-form-item>
        <el-input v-model="loginForm.FormData.username" placeholder="请输入用户名">
          <template #prefix>
            <i-ep-user />
          </template>
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-input
          v-model="loginForm.FormData.password"
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
            <el-checkbox v-model="loginForm.FormData.rememberMe" label="记住我"></el-checkbox>
          </el-col>
          <el-col :offset="6" :span="8">
            <el-link style="float: right" type="primary">忘记密码？</el-link>
          </el-col>
        </el-row>
      </el-form-item>

      <el-form-item>
        <el-button :loading="loginLoading" style="width: 100%;" type="primary" @click="handleLogin">登陆</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import * as UserApi from '@/api/user'
import * as AuthUtils from '@/utils/auth'
import router from '@/router'


const loginForm = reactive({
  FormData: {
    username: 'admin',
    password: '123456',
    rememberMe: undefined
  }
});

// 加载状态
const loginLoading = ref(false);

// 登录方法
const handleLogin = async () => {
  loginLoading.value = true;

  try{
    const res = await UserApi.login(loginForm.FormData)
    if (!res || res.code !== 200) {
      return
    }
    AuthUtils.setToken(res.data.token)

    // 跳转到首页
    await router.push('/index')

  } finally {
    loginLoading.value = false;
  }

}


</script>

<style scoped></style>

