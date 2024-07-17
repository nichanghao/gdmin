import { createApp } from 'vue';
import { setupStore } from '@/stores';
// 路由
import router, { setupRouter } from '@/router';
// 确保在创建应用实例之前导入了导航守卫
import '@/router/permission';
import App from './App.vue';

const app = createApp(App);

// 使用 pinia 状态管理
setupStore(app);

// 设置路由
setupRouter(app);
await router.isReady();

app.mount('#app');
