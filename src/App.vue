<template>
	<div id="app">
		<div id="main">
			<Film />
		</div>
		<Footer />
	</div>
</template>

<script>
import Film from "./components/Film.vue";
import Footer from "./components/Footer.vue";
export default {
	name: "App",
	components: {
		Film,
		Footer,
	},
	mounted() {
		var pref = window.matchMedia("(prefers-color-scheme: dark)");
		if (pref.matches) {
			document.body.classList.add("dark");
		} else {
			document.body.classList.remove("dark");
		}
		window.matchMedia("(prefers-color-scheme: dark)").addListener((e) => {
			if (e.matches) {
				document.body.classList.add("dark");
			} else {
				document.body.classList.remove("dark");
			}
		});
	},
};
</script>

<style>
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
