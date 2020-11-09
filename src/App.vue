<template>
	<div id="app">
		<DarkModeToggle />
		<div id="main">
			<Film />
		</div>
		<Footer />
	</div>
</template>

<script>
import Film from "./components/Film.vue";
import Footer from "./components/Footer.vue";
import DarkModeToggle from "./components/DarkModeToggle.vue";
export default {
	name: "App",
	components: {
		Film,
		Footer,
		DarkModeToggle,
	},
	mounted() {
		var pref = window.matchMedia("(prefers-color-scheme: dark)");
		let darkModeIcon = document.getElementById("darkmode-icon");
		if (pref.matches) {
			document.body.classList.add("dark");
			darkModeIcon.classList.remove("moon");
		} else {
			document.body.classList.remove("dark");
			darkModeIcon.classList.add("moon");
		}
		window.matchMedia("(prefers-color-scheme: dark)").addListener((e) => {
			if (e.matches) {
				document.body.classList.add("dark");
				darkModeIcon.classList.remove("moon");
			} else {
				document.body.classList.remove("dark");
				darkModeIcon.classList.add("moon");
			}
		});
	},
};
</script>

<style>
/* #toggle {
	position: fixed;
	top: 0;
	right: 0;
	z-index: 999;
} */
#app {
	font-family: Avenir, Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
	margin-top: 30px;
}

#main {
	min-height: -webkit-calc(85vh - 15px);
}

body.dark {
	background: #1d232a;
}

.dark #app {
	color: #76a0ca;
}

::selection {
	background: #40bcf4;
	color: white;
}
</style>
