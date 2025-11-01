import {MovieType, movieTypeFromString} from './movie-types'

/**
 * A watchlist picker query
 */
export interface Query {
	/**
	 * list IDs; primarily user IDs (for watchlists), but can include actors
	 * or other non-watchlist lists by ID.
	 */
	lists: string[]

	/**
	 * Whether or not the lists should be intersected (where films picked from
	 * are present on all included lists).
	 */
	intersect: boolean

	/**
	 * Categories of films to ignore.
	 */
	ignore: MovieType[]

	/**
	 * Non-user visible option to perform the search server-side, returning
	 * a static page that doesn't require Javascript.
	 */
	no_javascript: boolean
}

/**
 * Key names in the `URLSearchParams` object
 */
export const key = {
	lists: 'u',
	intersect: 'i',
	ignore: 'ignore',
	no_javascript: 'nojs',
}

export function parseRequest(url: URL): Query {
	const request = {
		lists: getAllParams(url, key.lists, NO_OP),
		ignore: getAllParams(url, key.ignore, movieTypeFromString),
		intersect: !!url.searchParams.get(key.intersect),
		no_javascript: !!url.searchParams.get(key.no_javascript),
	}
	logRequest(request)
	return request
}

/**
 * Gets all values of the given search param key and applies a mapping function.
 */
function getAllParams<T>(url: URL, key: string, mapFunction: (param: string) => T): T[] {
	
	const params = url.searchParams.getAll(key) || []

	// For backwards compatibility with bookmarked queries. Params may be in
	// individual key-value pairs, comma separated lists, or space separated.
	return params.flatMap(it => it.split(/[\,\s]/g))
		.map(mapFunction)
		.filter(it => !!it) // remove empty values from extra commas/whitespace
}

const NO_OP = (value: string) => value

function logRequest(query: Query) {
	const timestamp = new Date().toISOString()
	const requestLog = `
		${timestamp}
		total_lists=${query.lists.length}
		${query.lists.map(id => `list_id=${id}`).join('\n') || 'list_id=[]'}
		intersection=${query.intersect}
		${query.ignore.map(opt => `ignore=${opt}`).join('\n') || 'ignore=[]'}
		no_javascript=${query.no_javascript}
	`
	console.log(requestLog.trim().replace(/\t/g,''))
}
