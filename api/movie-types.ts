
export enum MovieType {
	unreleased = 'unreleased',
	shorts = 'shorts',
	feature = 'feature',
}

/**
 * Converts a string value to `MovieType` member. Unknown values return
 * `undefined`. Note: supports "true" for legacy reasons.
 */
export function movieTypeFromString(input: string): MovieType | undefined {
	if (Object.values(MovieType).includes(input as MovieType)) {
		return input as MovieType
	}
	else if (input == 'true') {
		// For backwards compatibility with old bookmarked requests. The
		// `ignore` query param used to be a boolean value that would just
		// ignore unreleased films. So `?ignore=true` is synonymous with
		// `MovieTypes.unreleased`.
		return MovieType.unreleased
	}
	else {
		console.error(`Unknown movie type in ignore parameter: ${input}`)
		return undefined
	}  
}
