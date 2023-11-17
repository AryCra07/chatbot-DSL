import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import requireTransform from 'vite-plugin-require-transform';
import { fileURLToPath, URL } from 'node:url';
import eslintPlugin from 'vite-plugin-eslint';
// https://vitejs.dev/config/

export default defineConfig({
  plugins: [
    vue(),
    requireTransform({
      fileRegex: /.ts$|.vue$/,
    }),
    // eslintPlugin({
    //   include: ['src/**/*.ts', 'src/**/*.vue', 'src/*.ts', 'src/*.vue'],
    // }),
  ],
  server: {
    open: true,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8848/',
        changeOrigin: true,
        rewrite(path) {
          return path.replace(/^\/api/, '');
        },
      },
    },
  },
  define: {
    'process.env': {},
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
    extensions: ['.vue', '.ts', '.js'],
  },
});
