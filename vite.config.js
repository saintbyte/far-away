/* Vercel build config */

export default {
    publicDir: 'assets',
    root: 'frontend/',
    build: {
        minify: 'esbuild',
        outDir: '../dist',
        rollupOptions: {
            input: {
                app: './main.html',
            },
            output: {
                inlineDynamicImports : true,
                entryFileNames: `static/js/[name].js`,
                chunkFileNames: `static/assets/[name].js`,
                assetFileNames: `static/assets/[name].[ext]`
            }
        },
    }
}