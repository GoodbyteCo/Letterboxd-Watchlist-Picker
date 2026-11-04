import vercel from '@astrojs/vercel'
import {defineConfig} from 'astro/config'

export default defineConfig({
	adapter: vercel(),
	output: 'server',
	// Preserve inline spacing from Astro 5 instead of Astro 7's JSX whitespace rules.
	compressHTML: true,
	security: {checkOrigin: false},
	srcDir: '.',
})
