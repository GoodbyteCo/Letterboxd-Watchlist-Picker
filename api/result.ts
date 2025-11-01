
export type Result = MovieResult | NoResult

interface MovieResult {
	status: 'success'
	title: string
	url: string
	img: string
}

interface NoResult {
	status: NoResultReason
	title: undefined
	url: undefined
	img: undefined
}

export enum NoResultReason {
	NO_INTERSECTION,
	ALL_IGNORED,
	TIMED_OUT,
	UNKNOWN,
}

export function isNoResult(result: Result): result is NoResult {
	return (result.status != 'success')
}
