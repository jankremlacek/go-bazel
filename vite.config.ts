import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tsconfigPaths from 'vite-tsconfig-paths'

// https://vitejs.dev/config/
export default defineConfig({
    build: {
        /*
        chunkSizeWarningLimit: 2048,
        manifest: true,
        rollupOptions: {
            // overwrite default .html entry
            input: 'src/app.ts',
            output: {
                manualChunks: {
                    amcharts: ['@amcharts/amcharts5'],
                },
            },
        },
        outDir: 'service/dist',
        */
    },
    plugins: [
        tsconfigPaths({
            extensions: ['.ts', '.tsx', '.js', '.jsx', '.vue'],
            loose: true,
        }),
        vue(),
    ],
    publicDir: './public',
})
