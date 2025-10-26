---
import ScreenreaderOnly from './screenreader-only.astro'

interface Props {
	darkmode: boolean
}

const {darkmode} = Astro.props
---

<label id="darkmode-switch">
	<ScreenreaderOnly>
		Darkmode:
	</ScreenreaderOnly>
	<input
		id="darkmode-toggle"
		type="checkbox"
		checked={darkmode}
	/>
	<svg
		xmlns="http://www.w3.org/2000/svg"
		xmlns:xlink="http://www.w3.org/1999/xlink"
		viewBox="0 0 796 796"
		role="presentation"
	>
		<defs>
			<mask id="moon-mask">
				<rect id="moon-mask-rect" width="100%" height="100%" fill="#fff"/>
				<circle id="moon-slice" r="268"/>
			</mask>
		</defs>
		<circle id="sun" r="268" transform="matrix(1 0 0 1 398.18 397.81)" mask="url(#moon-mask)"/>
		<rect class="ray" width="30" height="101" rx="15" ry="15" transform="matrix(0.88 0.46 -0.46 0.88 258.93 627.58)"/>
		<rect class="ray" width="30" height="101" rx="15" ry="15" transform="matrix(0.95 -0.29 0.29 0.95 462.01 658.67)"/>
		<rect class="ray" width="30" height="101" rx="15" ry="15" transform="matrix(0.29 0.95 -0.95 0.29 137.32 462.01)"/>
		<rect class="ray" width="30" height="101" rx="15" ry="15" transform="matrix(-0.46 0.88 -0.88 -0.46 168.41 258.93)"/>
		<rect class="ray" width="30" height="101" rx="15" ry="15" transform="matrix(0.88 0.46 -0.46 0.88 557.99 65.15)"/>
		<rect class="ray" width="30" height="101" rx="15" ry="15" transform="matrix(0.95 -0.29 0.29 0.95 275.76 49.50)"/>
		<rect class="ray" width="30" height="101" rx="15" ry="15" transform="matrix(0.29 0.95 -0.95 0.29 746.49 275.76)"/>
		<rect class="ray" width="30" height="101" rx="15" ry="15" transform="matrix(-0.46 0.88 -0.88 -0.46 730.84 557.99)"/>
		<clipPath><rect width="796" height="796" rx="0" ry="0" fill="#000"/></clipPath>
	</svg>
</label>
<noscript>
	<style>
		#darkmode-switch
		{
			display: none;
		}
	</style>
</noscript>

<script>
	/** id for local storage and cookie */
	const DARK_MODE = 'dark-mode'

	/** sets a cookie with a 5-year expiry */
	function setCookie(key: string, value: string) {
		const expiry = new Date()
		expiry.setFullYear(expiry.getFullYear() + 5)
		document.cookie = `${encodeURIComponent(key)}=${encodeURIComponent(value)}; expires=${expiry.toUTCString()}; path=/`
	}

	/** gets the cookie of the given key; `undefined` if not found */
	function getCookie(key: string): string | undefined {
		const name = `${encodeURIComponent(key)}=`
		for (const cookie of document.cookie.split(';')) {
			const trimmed = cookie.trim()
			if (trimmed.startsWith(name)) {
				return decodeURIComponent(trimmed.substring(name.length))
			}
		}
		return undefined
	}

	/**
	 * sets the theme according to the current state of the checkbox. sets the
	 * value in local storage as well as cookies. then re-runs `applyTheme()`.
	 */
	function setTheme() {
		const toggle = document.getElementById('darkmode-toggle') as HTMLInputElement
		setCookie(DARK_MODE, (toggle.checked) ? 'true' : 'false')
		localStorage.setItem(DARK_MODE, (toggle.checked) ? '1' : '0')
		applyTheme()
	}

	function applyTheme() {

		const toggle = document.getElementById('darkmode-toggle') as HTMLInputElement
		const system = window.matchMedia('(prefers-color-scheme: dark)')
		const stored = localStorage.getItem(DARK_MODE) || undefined
		const cookie = getCookie(DARK_MODE)
		let darkmode = false

		// in the past we've only stored dark mode preferences to local storage,
		// but this requires client-side JS to run causing a flashbang on load.
		// cookies let us set the theme server-side and send the markup over with
		// dark mode already set. so if we don't have a cookie yet, let's make one!
		if (!cookie) {
			// dark mode if they've stored a preference for it, or if their
			// system specifies it but they haven't saved a preference. 
			darkmode = (stored == '1' || (system.matches && stored == undefined))
			setCookie(DARK_MODE, `${darkmode}`)
		}
		else {
			// once a cookie is set, that's all we care about.
			darkmode = (cookie == 'true')
		}
		
		document.body.classList.toggle('dark', darkmode)
		document.querySelector('meta[name=theme-color]')?.setAttribute('content', darkmode ? '#14181d' : '#fff')
		toggle.checked = darkmode
    }

	window.addEventListener('load', () => {
		applyTheme()
		document.getElementById('darkmode-toggle')?.addEventListener('change', setTheme)
	})
</script>

<style>
	label
	{
		position: absolute;
		inset-block-start: 15px;
		inset-inline-end: 15px;

		width: 40px;
		height: 20px;
		border-radius: 20px;

		color: transparent;
		user-select: none;
		line-height: 30px;

		--off-white: #c2c2c2; /* darken for better contrast */
		background-color: var(--off-white);
		transition: background-color 0.3s ease;
		cursor: pointer;
	}

	label:has(input:checked)
	{
		background-color: var(--tertiary);
	}

	label:focus,
	label:has(input:focus-visible)
	{
		box-shadow:
			0 0 0 3px var(--background),
			0 0 0 5px #275ec5;
	}

	label::after
	{
		content: "";
		position: absolute;

		inset-block-start: 1px;
		inset-inline: 21px 1px;
		width: 18px;
		height: 18px;

		border-radius: 50%;
		background-color: var(--white);
		transition: all 0.3s;
	}

	label:has(input:checked)::after
	{
		inset-inline: 1px 21px;
	}

	input
	{
		opacity: 0;
		outline: none;
		width: 0px;
		height: 0px;
	}

	svg
	{
		width: 30px;
		position: absolute;
		inset-block-start: -5px;
		inset-inline-start: -40px;
		fill: var(--off-white);
	}

	:global(body.dark) svg
	{
		fill: var(--tertiary);
	}

	svg #sun
	{
		transition: r 0.3s ease;
		r: 202px;
	}

	input:checked + svg #sun
	{
		r: 268px;
	}

	svg .ray 
	{
		transition: opacity 0.3s ease;
		opacity: 1;
	}

	input:checked + svg .ray
	{
		opacity: 0;
	}

	svg #moon-slice
	{
		transition: fill 0.3s ease, transform 0.3s;
		fill: #000; /* color for svg clipping */
		transform: translate(560.796528px, -311.186989px);
	}

	input:checked + svg #moon-slice
	{
		transform: translate(220px, -120px);
	}

	svg #moon-mask-rect
	{
		transform: translate(-268px, -268px);
		fill: #fff; /* color for svg clipping */
	}
</style>
