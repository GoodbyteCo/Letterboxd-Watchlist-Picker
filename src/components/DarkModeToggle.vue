<template>
	<div 
		v-on:keyup.enter="toggleDarkModeOn"
		id="darkmode-toggle-box"
		tabindex="1"
	>
		<svg
			v-on:click="toggleDarkModeOn"
			id="darkmode-icon"
			xmlns="http://www.w3.org/2000/svg"
			xmlns:xlink="http://www.w3.org/1999/xlink"
			viewBox="0 0 796 796"
			shape-rendering="geometricPrecision"
			text-rendering="geometricPrecision"
		>
			<defs>
				<mask id="moon-mask">
					<rect
						id="moon-mask-rect"
						width="100%"
						height="100%"
						fill="#fff"
					></rect>
					<circle
						id="moon-slice"
						r="268"
						stroke="none"
						stroke-width="1"
					/>
				</mask>
			</defs>

			<circle
				id="sun-body"
				r="268"
				transform="matrix(1 0 0 1 398.18936602511235 397.81063397507751)"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
				mask="url(#moon-mask)"
			/>

			<rect
				class="sun-beam"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.88300001621246 0.46900001168251 -0.46900001168251 0.88300001621246 258.93798828125000 627.58801269531250)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				class="sun-beam"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.95630475596304 -0.29237170472274 0.29237170472274 0.95630475596304 462.01098632812500 658.67498779296875)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				class="sun-beam"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.29237170472274 0.95630475596304 -0.95630475596304 0.29237170472274 137.32499694824219 462.01098632812500)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				class="sun-beam"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(-0.46900001168251 0.88300001621246 -0.88300001621246 -0.46900001168251 168.41200256347656 258.93701171875000)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				class="sun-beam"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.88300001621246 0.46900001168251 -0.46900001168251 0.88300001621246 557.99102783203125 65.15000152587891)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				class="sun-beam"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.95630475596304 -0.29237170472274 0.29237170472274 0.95630475596304 275.76998901367188 49.50899887084961)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				class="sun-beam"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.29237170472274 0.95630475596304 -0.95630475596304 0.29237170472274 746.49102783203125 275.76998901367188)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				class="sun-beam"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(-0.46900001168251 0.88300001621246 -0.88300001621246 -0.46900001168251 730.84997558593750 557.99102783203125)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<clipPath id="eg5pwu10v5713">
				<rect
					id="eg5pwu10v5714"
					width="796"
					height="796"
					rx="0"
					ry="0"
					fill="rgb(0,0,0)"
					stroke="none"
					stroke-width="1"
				/>
			</clipPath>
		</svg>

		<input
			v-model="darkmodeOn"
			type="checkbox"
			id="toggle"
			class="checkbox"
			v-on:change="swapdark"
		/>
		<label for="toggle" class="switch"></label>
	</div>
</template>
<script>
export default {
	name: "DarkModeToggle",
	data() {
		return {
			darkmodeOn: true,
		};
	},
	mounted() {
		var pref = window.matchMedia("(prefers-color-scheme: dark)");
		let darkModeIcon = document.getElementById("darkmode-icon");
		if (
			(pref.matches && localStorage.getItem("dark-mode") === null) ||
			localStorage.getItem("dark-mode") == 1
		) {
			document.body.classList.add("dark");
			darkModeIcon.classList.remove("moon");
			this.darkmodeOn = true;
		} else {
			document.body.classList.remove("dark");
			darkModeIcon.classList.add("moon");
			this.darkmodeOn = false;
			console.log(this.darkmodeOn);
		}
		window.matchMedia("(prefers-color-scheme: dark)").addListener((e) => {
			if (localStorage.getItem("dark-mode") === null) {
				console.log("woot");
				if (e.matches) {
					document.body.classList.add("dark");
					darkModeIcon.classList.remove("moon");
					this.darkmodeOn = true;
				} else {
					document.body.classList.remove("dark");
					darkModeIcon.classList.add("moon");
					this.darkmodeOn = false;
				}
			}
		});
	},
	methods: {
		swapdark() {
			let darkModeIcon = document.getElementById("darkmode-icon");
			if (this.darkmodeOn) {
				document.body.classList.add("dark");
				darkModeIcon.classList.remove("moon");
				localStorage.setItem("dark-mode", 1);
			} else {
				document.body.classList.remove("dark");
				darkModeIcon.classList.add("moon");
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
#darkmode-toggle-box {
	position: fixed;
	top: 5px;
	right: 0px;
	outline: none;
	padding: 0 10px 5px 10px;
	border-radius: 3px;
}

#darkmode-toggle-box:focus-visible {
	box-shadow: 0 0 0 3px #fff, 0 0 0 5px #40bcf4;
}

.dark #darkmode-toggle-box:focus-visible {
	box-shadow: 0 0 0 3px #14181d, 0 0 0 5px #40bcf4;
}

#darkmode-icon {
	width: 30px;
	transform: translateY(5px);
	margin-right: 10px;
}

#darkmode-icon * {
	fill: #c2c2c2;
}

.dark #darkmode-icon * {
	fill: #526e89;
}

.switch {
	position: relative;
	display: inline-block;
	width: 40px;
	height: 20px;
	background-color: #c2c2c2;
	border-radius: 20px;
	transition: background-color 0.3s ease;
}

.switch::after {
	content: "";
	position: absolute;
	width: 18px;
	height: 18px;
	border-radius: 50%;
	background-color: white;
	top: 1px;
	left: 1px;
	transition: all 0.3s;
}

.checkbox:checked + .switch::after {
	left: 20px;
}

.checkbox:checked + .switch {
	background-color: #526e89;
}

.checkbox {
	display: none;
}

/* icon animation  */

#sun-body {
	transition: all 0.3s ease;
	r: 202px;
}

.moon #sun-body {
	r: 268px;
}

.dark #darkmode-icon #moon-slice,
#darkmode-icon #moon-slice {
	transition: all 0.3s ease;
	fill: #000;
	transform: translate(560.796528px, -311.186989px);
}

.dark #darkmode-icon.moon #moon-slice,
#darkmode-icon.moon #moon-slice {
	transform: translate(220px, -120px);
}

.dark #darkmode-icon #moon-mask-rect,
#darkmode-icon #moon-mask-rect {
	transform: translate(-268px, -268px);
	fill: white;
}

.sun-beam {
	transition: all 0.3s ease;
	opacity: 1;
}

.moon .sun-beam {
	opacity: 0;
}
</style>
