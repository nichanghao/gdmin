// element-plus 按需导入
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import IconsResolver from 'unplugin-icons/resolver'
import Icons from 'unplugin-icons/vite'

import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'


const pathSrc = path.resolve(__dirname, 'src')

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        // element-plus 按需导入
        AutoImport({
            // 自动导入 Element Plus 相关函数，如：ElMessage, ElMessageBox... (带样式)
            resolvers: [
                ElementPlusResolver(),

                // 自动导入图标组件
                IconsResolver({
                prefix: 'Icon',
                }),
            ],
            dts: path.resolve(pathSrc, 'auto-imports.d.ts'),
        }),
        Components({
            resolvers: [
                // 自动注册图标组件
                IconsResolver({
                    enabledCollections: ['ep'],
                }),
                ElementPlusResolver()
            ],
            dts: path.resolve(pathSrc, 'components.d.ts'),
        }),
        Icons({
            autoInstall: true,
        }),
    ],
    resolve: {
        alias: {
            '@': pathSrc
        }
    }
})
