import { createApp } from 'vue'
import { setupStore } from '@/stores'
import App from './App.vue'
import router from './router'

const app = createApp(App)

// 使用 pinia 状态管理
setupStore(app)

app.use(router)

app.mount('#app')
