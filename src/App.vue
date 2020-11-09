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
	min-height: calc(85vh - 15px);
	min-height: -webkit-calc(85vh - 15px);
	min-height: -moz-calc(85vh - 15px);
}

body.dark {
	background: #14181d;
}

.dark #app {
	color: #526e89;
}

::selection {
	background: #40bcf4;
	color: white;
}
</style>
