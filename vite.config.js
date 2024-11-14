// vite.config.js
import { resolve } from 'path'
import { defineConfig } from 'vite'

export default defineConfig({
    build: {
        editor: {
            entry: resolve(__dirname, 'frontend/editor.js'),
            name: 'Editor',
            // the proper extensions will be added
            fileName: '/js/editor.js',
        },
        rollupOptions: {
            input: {
              edit: resolve(__dirname, 'index.html'),
            },
        },
    },
})