import tsconfigPaths from 'vite-tsconfig-paths';
import react from '@vitejs/plugin-react-swc';
import { defineConfig } from 'vitest/config';
import svgr from 'vite-plugin-svgr';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [tsconfigPaths({ root: '.' }), react(), svgr()],
  base: '/pmm-ui',
  resolve: {
    dedupe: [
      'react',
      'react-dom',
      '@emotion/react',
      '@emotion/styled',
      '@mui/material',
      '@mui/system',
      '@mui/styled-engine',
    ],
  },
  optimizeDeps: {
    // Uncomment when using yarn link for @percona/percona-ui locally
    //exclude: ['@percona/percona-ui'],
    force: true,
  },
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          'vendor-react': ['react', 'react-dom', 'react-router-dom'],
          'vendor-emotion': ['@emotion/react', '@emotion/styled'],
          'vendor-mui-icons': ['@mui/icons-material'],
          'vendor-mui-material': ['@mui/material'],
          'vendor-mui-date-pickers': ['@mui/x-date-pickers'],
          'vendor-percona-ui': ['@percona/percona-ui'],
          'vendor-markdown': [
            'react-markdown',
            'react-syntax-highlighter',
            'rehype-raw',
            'remark-gfm',
          ],
          'vendor-date': ['@date-fns/tz', 'date-fns'],
        },
      },
    },
  },
  server: {
    watch: {
      // Watch the linked package for changes (negated pattern means "don't ignore")
      ignored: ['!**/node_modules/@percona/percona-ui/**'],
    },
    proxy: {
      '/v1': {
        target: '/',
      },
    },
    host: '0.0.0.0',
    strictPort: true,
    hmr: {
      clientPort: 5173,
    },
    allowedHosts: ['host.docker.internal'],
  },
  test: {
    globals: true,
    environment: 'jsdom',
    setupFiles: 'src/setupTests.ts',
    server: {
      deps: {
        fallbackCJS: true,
      },
    },
  },
});
