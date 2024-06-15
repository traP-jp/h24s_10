import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vuetify from "vite-plugin-vuetify";
import { VitePWA } from "vite-plugin-pwa";

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: [{ find: "/@/", replacement: "/src/" }],
  },
  plugins: [
    vue(),
    vuetify(),
    VitePWA({
      registerType: "autoUpdate",
      workbox: {
        globPatterns: ["**/*.{js,css,html,ico,png,svg}"],
        runtimeCaching: [
          {
            urlPattern: /^https:\/\/q.trap.jp\/api\/v3\/public\/icon\/.*/i,
            handler: "StaleWhileRevalidate",
            options: {
              cacheName: "traQ-icons",
              expiration: {
                maxEntries: 10,
                maxAgeSeconds: 60 * 60 * 24 * 365, // <== 365 days
              },
              cacheableResponse: {
                statuses: [0, 200],
              },
            },
          },
        ],
      },
      devOptions: {
        enabled: true,
      },
    }),
  ],
  server: {
    proxy: {
      "/api": {
        target: "https://h24s10.trap.show/api/",
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, "/"),
      },
    },
  },
});
