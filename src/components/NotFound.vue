---
import {NoResultReason} from '../../api/result'

interface Props {
	reason: NoResultReason
}

function explaination(reason: NoResultReason) {
	switch (reason) {
		case NoResultReason.ALL_IGNORED: return `
			There were no films found in that list. It may be
			empty, private, or only contain films ignored by
			settings (check Advanced Options above).
		`
		case NoResultReason.NO_INTERSECTION: return `
			The intersection between those two lists is empty.
		`
		case NoResultReason.TIMED_OUT: return `
			Sorry your list was too powerful and we timed out. Try 
			another list.
		`
		default: return `
			Sorry, that watchlist is empty, private, or doesn't
			exist at all.`
	}
}

const {reason} = Astro.props
---

<h2>Nothing Found</h2>
<p>{explaination(reason)}</p>
<img
	id="not-found-gif"
	width="250"
	src="/poe.gif"
	alt="from the movie Kung-Fu Panda, protagonist Poe looks down at an empty scroll"
/>

<style>
	h2
	{
		font-size: 1.2rem;
	}

	p
	{
		max-width: 60ch;
		margin: 0rem auto 1rem;
		padding: 0 2rem;
	}

	img
	{
		margin: 30px 0;
		border-radius: 4px;
	}
</style>
