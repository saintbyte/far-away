/* Vercel build config */
import { terser } from 'rollup-plugin-terser'

export default {
    publicDir: 'assets',
    root: 'frontend/',
    build: {
        minify: 'terser',
        outDir: '../dist',
        rollupOptions: {
            plugins: [
                terser() // Use terser plugin
            ],
            output: {
                entryFileNames: `js/[name].js`,
                chunkFileNames: `assets/[name].js`,
                assetFileNames: `assets/[name].[ext]`
            }
        },
    }
}