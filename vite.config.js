/* Vercel build config */

export default {
    publicDir: 'assets',
    root: 'frontend/',
    build: {
        minify: 'esbuild',
        outDir: '../dist',
        rollupOptions: {
            output: {
                inlineDynamicImports : true,
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