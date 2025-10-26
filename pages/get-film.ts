// @ts-ignore
import {experimental_AstroContainer} from 'astro/container'
// @ts-ignore
import ServerSideResults from '../components/result/server-side.astro'

/**
 * Takes a request in the form of a `Query`, and returns the rendered HTML
 * of that query's `ServerSideResults`. This allows results to be rendered
 * on the server and passed back to the client to be injected into the site.
 * Note: uses experimental features, may break on Astro version bumps!
 * See: https://docs.astro.build/en/reference/container-reference/
 */
export async function POST({request}) {
	try {
		const query = await request.json()
		const astro = await experimental_AstroContainer.create({streaming: false})
		const html = await astro.renderToString(ServerSideResults, {props: {query}})
		return new Response(html, {
			status: 200,
			headers: {'Content-Type': 'text/html; charset=utf-8'},
		})
	}
	catch (error) {
		console.error('Error building client-side results', error)
		return new Response(JSON.stringify({error: error.message}), {
			status: 500,
			headers: {'Content-Type': 'application/json'},
		})
	}
}
