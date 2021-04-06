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
};
</script>

<style>
	@media (prefers-reduced-motion) 
	{
		* 
		{
			animation: none !important;
			transition: none !important;
		}
	}

	:root
	{
		--background: #fff;
		--foreground: #2c3e50;
		--primary: #1caff2;
		--secondary: #000;
		--tertiary: #fff;
	}

	.dark
	{
		--background: #14181d;
		--foreground: #76a0ca; /* #778fa6; */
		--primary: #1caff2;
		--secondary: #fff;
		--tertiary: #fff;
	}

	body 
	{
		background: var(--background);
		transition: background ease-in-out 0.25s;
	}

	#app 
	{
		font-family: Avenir, Helvetica, Arial, sans-serif;
		color: var(--foreground);
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		text-align: center;

		transition: color ease-in-out 0.25s;
		margin-top: 30px;
	}

	main 
	{
		min-height: calc(85vh - 15px);
		min-height: -webkit-calc(85vh - 15px);
		min-height: -moz-calc(85vh - 15px);

		transform: translateY(40px);
		transition: transform 1.2s cubic-bezier(0.82, 0.01, 0.45, 1);
	}

	.entered main,
	.done main
	{
		transform: none;
	}

	::selection 
	{
		background: var(--primary);
		color: var(--tertiary);
	}
</style>
