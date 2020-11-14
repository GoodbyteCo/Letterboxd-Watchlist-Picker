<template>
	<div id="app">
		<DarkModeToggle />
		Hello
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
};
</script>

<style>
@media (prefers-reduced-motion) {
	* {
		animation: none !important;
		transition: none !important;
	}
}
/* #toggle {
	position: fixed;
	top: 0;
	right: 0;
	z-index: 999;
} */

body {
	background: #fff;
	transition: background ease-in-out 0.25s;
}
#app {
	font-family: Avenir, Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
	transition: color ease-in-out 0.25s;
	margin-top: 30px;
}

#main {
	min-height: calc(85vh - 15px);
	min-height: -webkit-calc(85vh - 15px);
	min-height: -moz-calc(85vh - 15px);
}

body.dark {
	background: #14181d;
	transition: background ease-in-out 0.25s;
}

.dark #app {
	color: #778fa6;
	transition: color ease-in-out 0.25s;
}

::selection {
	background: #40bcf4;
	color: white;
}
</style>
