/* Build to local public dir */

export default {

    publicDir: 'assets',
    root: 'frontend/',
    build: {
        outDir: '../public',
        minify: 'terser',
        rollupOptions: {
            output: {
                inlineDynamicImports : true,
                entryFileNames: `js/[name].js`,
                chunkFileNames: `assets/[name].js`,
                assetFileNames: `assets/[name].[ext]`
            }
        },
    },
    server: {
        proxy: {
            // string shorthand: http://localhost:5173/api -> http://localhost:8080/api
            '/api': 'http://localhost:8080',
        }
    }
}