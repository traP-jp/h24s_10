import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: [{ find: '/@/', replacement: '/src/' }]
  },
  plugins: [vue()],
  server: {
    proxy: {
      '/api': {
        target: 'https://exercise-app-kari-backend.trap.games',
        changeOrigin: true,
        rewrite: path => path.replace(/^\/api/, '/')
      }
    }
  }
})
