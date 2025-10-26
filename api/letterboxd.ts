import {Query} from './query'
import {NoResultReason, Result} from './result'

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

		const html = await response.text()
		// todo: parse?

		return {
			status: 'success',
			title: html,
			img:'https://a.ltrbxd.com/resized/sm/upload/8g/5p/p4/6b/8YWirGQidtZeSEmhqvQM5FrI6N1-0-230-0-345-crop.jpg?v=10e3695a0f',
			url: letterboxdUrl.toString(),
		}
	}
	catch (error) {
		console.error(`An error occurred:`, error)
		return buildError(NoResultReason.TIMED_OUT)
	}	
}

/**
 * Gets a random item from the given list.
 */
function pickRandom(options: string[]) {
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
