/* Vercel build config */
export default {
    publicDir: 'assets',
    root: 'frontend/',
    build: {
        outDir: 'public',
        file: 'main.js',
        format: 'iife',
        name: 'MyBundle'
    }
}