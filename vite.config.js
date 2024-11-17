/* Vercel build config */
export default {
    publicDir: 'assets',
    root: 'frontend/',
    build: {
        outDir: '../dist',
        file: 'main.js',
        format: 'iife',
        name: 'MyBundle'
    }
}