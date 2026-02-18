import { defineConfig } from 'vitest/config'
import vue from '@vitejs/plugin-vue'

// Reuse the same proxy logic if needed for tests, or just basic setup
export default defineConfig({
    plugins: [vue()],
    test: {
        globals: true,
        environment: 'jsdom',
        alias: {
            '@': '/src'
        }
    },
})
