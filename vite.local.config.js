/* Build to local public dir */
import {terser} from "rollup-plugin-terser";

export default {

    publicDir: 'assets',
    root: 'frontend/',
    build: {
        outDir: '../public',
        minify: 'terser',
        rollupOptions: {
            plugins: [
                terser() // Use terser plugin
            ],
            output: {
                inlineDynamicImports : true,
                entryFileNames: `js/[name].js`,
                chunkFileNames: `assets/[name].js`,
                assetFileNames: `assets/[name].[ext]`
            }
        },

    }
}