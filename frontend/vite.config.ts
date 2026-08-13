import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: true,
    port: 5173,
    proxy: {
      "/host/report": "http://go:4000/",
      "/file/report": "http://go:4000/",
    },
    watch: {
      usePolling: true,
    },
    hmr: {
      port: 5173,
    }
  },
})
