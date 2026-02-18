import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// API target: use env var in Docker, fallback to localhost for local dev
const apiTarget = process.env.VITE_API_URL || 'http://localhost:3000'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    host: true, // allow access from Docker network
    proxy: {
      '/api': {
        target: apiTarget,
        changeOrigin: true,
      },
    },
  },
})
