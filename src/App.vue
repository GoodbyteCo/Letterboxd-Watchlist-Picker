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
		var pref = window.matchMedia("(prefers-color-scheme: dark)");
		if (pref.matches) {
			this.isActive = true;
		} else {
			this.isActive = false;
		}
		window.matchMedia("(prefers-color-scheme: dark)").addListener((e) => {
			if (e.matches) {
				this.isActive = true;
			} else {
				this.isActive = false;
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

/* @media (prefers-color-scheme: dark) { */
.dark body {
	background: #1d232a;
}
.dark #app {
	color: #76a0ca;
}
/* } */

::selection {
	background: #40bcf4;
	color: white;
}
</style>
