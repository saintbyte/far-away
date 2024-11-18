/* Vercel build config */
import { terser } from 'rollup-plugin-terser'

export default {
    publicDir: 'assets',
    root: 'frontend/',
    build: {
        minify: 'terser',
        outDir: '../dist',
        file: 'main.js',
        format: 'iife',
        name: 'MyBundle',
        rollupOptions: {
            plugins: [
                terser() // Use terser plugin
            ],
            input: 'editor.js', // Your entry point
            output: {
                file: 'editor.min.js', // Output file
                format: 'es', // Output format
                name: 'Editor' // Name of the global variable
            }
        }
    }
}