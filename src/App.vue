<template>
	<div>
		<header>
			<logo/>
			<dark-mode-toggle/>
		</header>
		<main>
			<about-text/>
			<search-bar v-model="users" :action="() => submit()">
				<advanced-options v-model="advancedOptions"/>
			</search-bar>
			<section id="film-results">
				<loading-bar v-if="loading"/>
			</section>
			<Film />
		</main>
		<Footer />
	</div>
</template>

<script>
	import Film from './components/Film.vue';
	import Footer from './components/Footer.vue';
	import DarkModeToggle from './components/DarkModeToggle.vue';
	import Logo from './components/Logo.vue';
	import AboutText from './components/AboutText.vue';
	import SearchBar from './components/SearchBar.vue';
	import AdvancedOptions from './components/AdvancedOptions.vue';
	import LoadingBar from './components/LoadingBar.vue';

	export default 
	{
		name: "App",
		components: 
		{
			Film,
			Footer,
			DarkModeToggle,
			Logo,
			AboutText,
			SearchBar,
			AdvancedOptions,
			LoadingBar
		},
		data()
		{
			return {
				users: '',
				advancedOptions: {},
			}
		},
		methods:
		{
			submit()
			{
				alert(this.advancedOptions['selectionMode'] + ' ' + this.advancedOptions['ignoreUnreleased']) 
			}
		}
	};
</script>

<style>
	:root
	{
		--background: #fff;
		--foreground: #2c3e50;

		--primary: #1caff2;
		--secondary: #000;
		--tertiary: rgb(64 188 244 / 0.5);

		--white: #fff;
		--off-white: #ebebeb;
		--black: #000;
	}

	.dark
	{
		--background: #14181d;
		--foreground: #76a0ca;

		--secondary: #fff;
		--tertiary: #526e89;
	}

	body 
	{
		background: var(--background);

		font-family: Avenir, Helvetica, Arial, sans-serif;
		color: var(--foreground);
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		text-align: center;

		transition: color ease-in-out 0.25s, background ease-in-out 0.25s;
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
		color: var(--white);
	}

	@media (prefers-reduced-motion) 
	{
		* 
		{
			animation: none !important;
			transition: none !important;
		}
	}
</style>
