import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { quasar, transformAssetUrls } from '@quasar/vite-plugin'
import wails from '@wailsio/runtime/plugins/vite'
import { join } from 'node:path' // 👈 Добавлен импорт для путей

export default defineConfig({
	plugins: [
		vue({
			template: { transformAssetUrls },
		}),
		quasar({
			sassVariables: join(__dirname, 'src/css/quasar.variables.scss'),
		}),
		wails('./bindings'),
	],
	server: {
		host: '127.0.0.1',
		port: Number(process.env.WAILS_VITE_PORT) || 9245,
		strictPort: true,
	},
})
