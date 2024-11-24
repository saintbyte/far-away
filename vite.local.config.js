/* Build to local public dir */

export default {

    publicDir: 'assets',
    root: 'frontend/',
    build: {
        outDir: '../public',
        minify: 'terser',
        rollupOptions: {
            input: {
                app: './main.html',
            },
            output: {
                inlineDynamicImports: true,
                entryFileNames: `static/js/[name].js`,
                chunkFileNames: `static/assets/[name].js`,
                assetFileNames: `static/assets/[name].[ext]`
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