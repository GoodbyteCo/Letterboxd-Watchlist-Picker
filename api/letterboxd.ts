import {Query} from './query'
import {NoResultReason, Result} from './result'
import * as cheerio from 'cheerio'

export async function fetchFilm(query: Query): Promise<Result> {
	const base = 'https://watchlistpicker.com/api'
	const userParam = `?users=${query.lists.join('&users=')}`
	const intersectParam = (query.intersect) ? '&intersect=true' : ''
	const ignoreParam = (query.ignore.length > 0)
		? `&ignore=${query.ignore.join(',')}`
		: ''

	const apiUrl = `${base}${userParam}${intersectParam}${ignoreParam}`
	return fetch(apiUrl)
		.then(res => res.json())
		.then(result => ({
			status: 'success',
			title: result['film_name'],
			url: result['slug'],
			img: result['image_url'],
		}))
}

/**
 * todo:
 * 	- support actor/director lookups
 * 	- support ignore options
 * 	- support intersection
 */
export async function findFilm(query: Query): Promise<Result> {
	if (query.lists.length == 0) {
		return buildError(NoResultReason.UNKNOWN)
	}

	// note: this trick only works if unioning the lists together
	const list = pickRandom(query.lists)

	const letterboxdUrl = (looksLikeUserList(list))
		? urlFromUserList(list)
		: urlFromUsername(list)

	try {
		const response = await fetch(letterboxdUrl)

		if (!response.ok) {
			console.error(`Error fetching ${letterboxdUrl}: ${response.status} ${response.statusText}`)
			return buildError(NoResultReason.UNKNOWN)
		}

		return await response.text()
			.then(cheerio.load) // parse into virtual DOM
			.then(extractFilmData)
			.then(buildResult)
	}
	catch (error) {
		console.error(`An error occurred:`, error)
		return buildError(NoResultReason.TIMED_OUT)
	}	
}

/**
 * Parses out the film data of the first film in the poster grid. Expects markup like:
 * ```
 * <div class="poster-grid">
 *     <ul>
 *         <li class="tooltip griditem" data-original-title="The Spirit of the Beehive (1973)">
 *             <div
 *                 class="react-component"
 *                 data-component-class="LazyPoster"
 *                 data-request-poster-metadata="true"
 *                 data-likeable="true"
 *                 data-watchable="true"
 *                 data-rateable="true"
 *                 data-image-width="150"
 *                 data-image-height="225"
 *                 data-item-name="The Spirit of the Beehive (1973)"
 *                 data-item-slug="the-spirit-of-the-beehive"
 *                 data-item-link="/film/the-spirit-of-the-beehive/"
 *                 data-item-full-display-name="The Spirit of the Beehive (1973)"
 *                 data-film-id="49296"
 *                 data-empty-poster-src="https://s.ltrbxd.com/static/img/empty-poster-150-DtnLDE3k.png"
 *                 data-is-linked="true"
 *                 data-target-link="/film/the-spirit-of-the-beehive/"
 *                 data-details-endpoint="/film/the-spirit-of-the-beehive/json/"
 *                 data-show-menu="true"
 *                 data-hide-tooltip="true"
 *             >
 *             </div>
 *         </li>
 *         ...
 *     </ul>
 * </div>
 * ```
 */
function extractFilmData(queryDocument: cheerio.CheerioAPI) {
	const data = queryDocument('.poster-grid li div').attr()
	return {
		slug: data['data-item-slug'],
		name: data['data-item-full-display-name'],
		id: data['data-film-id'],
	}
}

function buildResult(film: ReturnType<typeof extractFilmData>): Result {
	return {
		status: 'success',
		title: film.name,
		url: `https://letterboxd.com/film/${film.slug}`,
		img: buildImageUrl(film),
	}
}

function buildImageUrl({slug, id}: ReturnType<typeof extractFilmData>) {
	// NOTE: there is a bug where sometimes the poster doesn't load. This is (often)
	// because the slug used in the film url differs slightly from the one used on
	// the image. For example:
	// - https://a.ltrbxd.com/resized/film-poster/8/9/8/0/6/6/898066-close-your-eyes-0-460-0-690-crop.jpg
	// - https://a.ltrbxd.com/resized/film-poster/8/9/8/0/6/6/898066-close-your-eyes-2023-0-460-0-690-crop.jpg
	//
	// The first one works (is Letterboxd's) but drops the `2023` from the films
	// slug. The second one is how we would generate the url and does not work.
	// I can't find any data on the page that reveals what this data is, and the
	// images are now all lazy loaded. Perhaps we could try cutting off the year
	// if the inital image 404s?
	return `https://a.ltrbxd.com/resized/film-poster/${id.split('').join('/')}/${id}-${slug}-0-460-0-690-crop.jpg`
}

/**
 * Gets a random item from the given list.
 */
function pickRandom<T>(options: T[]): T {
	if (options.length == 1) {
		return options[0]
	}
	return options[Math.floor(Math.random() * options.length)]
}

/**
 * Whether the given list entry looks like a user-created list (rather than
 * a watchlist). Like `qjack/my-list`, etc.
 */
function looksLikeUserList(list: string) {
	return list.includes('/')
}

function urlFromUsername(user: string): URL {
	return new URL(`https://letterboxd.com/${user}/watchlist/by/shuffle/`)
}

function urlFromUserList(listId: string): URL {
	// allows you to enter `qjack/my-list` or `qjack/list/my-list`
	const slug = listId.includes('/list/')
		? listId
		: listId.replace('/', '/list/')

	return new URL(`https://letterboxd.com/${slug}/by/shuffle/`)
}

function buildError(reason: NoResultReason): Result {
	return {
		status: reason,
		title: undefined,
		url: undefined,
		img: undefined,
	}
}
