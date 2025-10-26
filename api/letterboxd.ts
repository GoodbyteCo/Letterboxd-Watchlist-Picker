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

