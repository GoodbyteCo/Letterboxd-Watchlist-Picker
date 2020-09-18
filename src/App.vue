<template>
	<div v-bind:class="{ dark: isActive }" id="app">
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
	data() {
		return {
			isActive: false,
		};
	},
	mounted() {
		var dark = "#1d232a";
		var white = "#fff";
		var pref = window.matchMedia("(prefers-color-scheme: dark)");
		if (pref.matches) {
			this.isActive = true;
			document.body.style.background = dark;
		} else {
			this.isActive = false;
			document.body.style.background = white;
		}
		window.matchMedia("(prefers-color-scheme: dark)").addListener((e) => {
			if (e.matches) {
				this.isActive = true;
				document.body.style.background = dark;
			} else {
				this.isActive = false;
				document.body.style.background = white;
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

.dark#app {
	color: #76a0ca;
}

::selection {
	background: #40bcf4;
	color: white;
}
</style>
