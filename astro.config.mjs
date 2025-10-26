import netlify from '@astrojs/netlify' // todo: replace with vercel?
import {defineConfig} from 'astro/config'

export default defineConfig({
	adapter: netlify(),
	output: 'server',
	security: {checkOrigin: false},
	srcDir: '.',
})
