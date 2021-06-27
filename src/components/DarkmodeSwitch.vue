<template>
	<div class="container">
		<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 796 796" shape-rendering="geometricPrecision" text-rendering="geometricPrecision"
			id="darkmode-icon"
			v-on:click="toggleDarkModeOn">
			<defs>
				<mask id="moon-mask">
					<rect id="moon-mask-rect" width="100%" height="100%" fill="#fff"/>
					<circle id="moon-slice" r="268" stroke="none" stroke-width="1"/>
				</mask>
			</defs>
			<circle id="sun-body" r="268" transform="matrix(1 0 0 1 398.18936602511235 397.81063397507751)" fill="rgb(139,139,139)" stroke="none" stroke-width="1" mask="url(#moon-mask)"/>
			<rect class="sun-beam" width="30" height="101" rx="15" ry="15" transform="matrix(0.88300001621246 0.46900001168251 -0.46900001168251 0.88300001621246 258.93798828125000 627.58801269531250)" opacity="0" fill="rgb(139,139,139)" stroke="none" stroke-width="1"/>
			<rect class="sun-beam" width="30" height="101" rx="15" ry="15" transform="matrix(0.95630475596304 -0.29237170472274 0.29237170472274 0.95630475596304 462.01098632812500 658.67498779296875)" opacity="0" fill="rgb(139,139,139)" stroke="none" stroke-width="1"/>
			<rect class="sun-beam" width="30" height="101" rx="15" ry="15" transform="matrix(0.29237170472274 0.95630475596304 -0.95630475596304 0.29237170472274 137.32499694824219 462.01098632812500)" opacity="0" fill="rgb(139,139,139)" stroke="none" stroke-width="1"/>
			<rect class="sun-beam" width="30" height="101" rx="15" ry="15" transform="matrix(-0.46900001168251 0.88300001621246 -0.88300001621246 -0.46900001168251 168.41200256347656 258.93701171875000)" opacity="0" fill="rgb(139,139,139)" stroke="none" stroke-width="1"/>
			<rect class="sun-beam" width="30" height="101" rx="15" ry="15" transform="matrix(0.88300001621246 0.46900001168251 -0.46900001168251 0.88300001621246 557.99102783203125 65.15000152587891)" opacity="0" fill="rgb(139,139,139)" stroke="none" stroke-width="1"/>
			<rect class="sun-beam" width="30" height="101" rx="15" ry="15" transform="matrix(0.95630475596304 -0.29237170472274 0.29237170472274 0.95630475596304 275.76998901367188 49.50899887084961)" opacity="0" fill="rgb(139,139,139)" stroke="none" stroke-width="1"/>
			<rect class="sun-beam" width="30" height="101" rx="15" ry="15" transform="matrix(0.29237170472274 0.95630475596304 -0.95630475596304 0.29237170472274 746.49102783203125 275.76998901367188)" opacity="0" fill="rgb(139,139,139)" stroke="none" stroke-width="1"/>
			<rect class="sun-beam" width="30" height="101" rx="15" ry="15" transform="matrix(-0.46900001168251 0.88300001621246 -0.88300001621246 -0.46900001168251 730.84997558593750 557.99102783203125)" opacity="0" fill="rgb(139,139,139)" stroke="none" stroke-width="1"/>
			<clipPath>
				<rect width="796" height="796" rx="0" ry="0" fill="rgb(0,0,0)" stroke="none" stroke-width="1"/>
			</clipPath>
		</svg>
		<input
			id="toggle"
			class="checkbox"
			type="checkbox"
			v-model="darkmodeOn"
			v-on:change="swapdark"
			v-on:keyup.enter="toggleDarkModeOn"
		/>
		<label for="toggle" class="switch">
			darkmode
		</label>
	</div>
</template>

<script>
	export default
	{
		name: 'DarkmodeSwitch',
		data()
		{
			return {
				darkmodeOn: true,
			};
		},
		mounted() 
		{
			let pref = window.matchMedia("(prefers-color-scheme: dark)");
			let darkModeIcon = document.getElementById("darkmode-icon");

			if (
				(pref.matches && localStorage.getItem("dark-mode") === null) ||
				localStorage.getItem("dark-mode") == 1
			) {
				document.body.classList.add("dark");
				darkModeIcon.classList.remove("moon");
				this.darkmodeOn = true;
				document.querySelector("meta[name=theme-color]").setAttribute("content", "#14181d");
			} else {
				document.body.classList.remove("dark");
				darkModeIcon.classList.add("moon");
				this.darkmodeOn = false;
				document.querySelector("meta[name=theme-color]").setAttribute("content", "#fff");
			}

			window.matchMedia("(prefers-color-scheme: dark)").addListener(e => {
				if (localStorage.getItem("dark-mode") === null) {
					if (e.matches) {
						document.body.classList.add("dark");
						darkModeIcon.classList.remove("moon");
						this.darkmodeOn = true;
						document.querySelector("meta[name=theme-color]").setAttribute("content", "#14181d");
					} else {
						document.body.classList.remove("dark");
						darkModeIcon.classList.add("moon");
						this.darkmodeOn = false;
						document.querySelector("meta[name=theme-color]").setAttribute("content", "#fff");
					}
				}
			});
		},
		methods: 
		{
			swapdark() 
			{
				let darkModeIcon = document.getElementById("darkmode-icon");

				if (this.darkmodeOn) {
					document.body.classList.add("dark");
					darkModeIcon.classList.remove("moon");
					document.querySelector("meta[name=theme-color]").setAttribute("content", "#14181d");
					localStorage.setItem("dark-mode", 1);
				} else {
					document.body.classList.remove("dark");
					darkModeIcon.classList.add("moon");
					document.querySelector("meta[name=theme-color]").setAttribute("content", "#fff");
					localStorage.setItem("dark-mode", 0);
				}
			},
			toggleDarkModeOn() {
				this.darkmodeOn = !this.darkmodeOn;
				this.swapdark();
			},
		},
	};
</script>

<style scoped>
	.container
	{
		--off-white: #c2c2c2; /* darken for better contrast */

		position: fixed;
		top: 5px;
		right: 5px;
		outline: none;
		padding: 0 10px 5px 10px;
		border-radius: 6px;
	}

	[v-focus-visible=true] .container:focus-within
	{
		box-shadow: 0 0 0 2px var(--primary);
	}

	label
	{
		position: relative;
		display: inline-block;
		width: 40px;
		height: 20px;
		border-radius: 20px;

		color: transparent;
		user-select: none;
		line-height: 30px;

		background-color: var(--off-white);
		transition: background-color 0.3s ease;
	}

	input:checked + label
	{
		background-color: var(--tertiary);
	}

	label::after
	{
		content: "";
		position: absolute;
		top: 1px;
		left: 1px;
		width: 18px;
		height: 18px;
		border-radius: 50%;
		background-color: var(--white);
		transition: all 0.3s;
	}

	input:checked + label::after
	{
		left: 20px;
	}

	input
	{
		opacity: 0;
		outline: none;
		width: 0px;
		height: 0px;
	}

	/* icon & animation  */

	svg
	{
		width: 30px;
		transform: translateY(5px);
	}

	svg *
	{
		fill: var(--off-white);
	}

	.dark svg *
	{
		fill: var(--tertiary);
	}

	#sun-body
	{
		transition: all 0.3s ease;
		r: 202px;
	}

	.moon #sun-body
	{
		r: 268px;
	}

	.sun-beam 
	{
		transition: all 0.3s ease;
		opacity: 1;
	}

	.moon .sun-beam 
	{
		opacity: 0;
	}

	#moon-slice
	{
		transition: all 0.3s ease;
		fill: #000000; /* color for svg clipping */
		transform: translate(560.796528px, -311.186989px);
	}

	.moon #moon-slice
	{
		transform: translate(220px, -120px);
	}

	#moon-mask-rect
	{
		transform: translate(-268px, -268px);
		fill: #ffffff; /* color for svg clipping */
	}
</style>
