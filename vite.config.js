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
                entryFileNames: `js/[name].js`,
                chunkFileNames: `assets/[name].js`,
                assetFileNames: `assets/[name].[ext]`
            }
        },
    }
}