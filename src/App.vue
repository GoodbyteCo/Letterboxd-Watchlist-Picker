<template>
	<div v-focus-visible>
		<header>
			<logo/>
			<darkmode-switch/>
		</header>
		<main>
			<about-text/>
			<search-bar v-model="users" :action="() => submit()">
				<advanced-options v-model="advancedOptions"/>
			</search-bar>
			<section id="film-results" aria-live="polite" :aria-busy="loading">
				<loading-bar v-if="loading"/>
				<div v-else-if="submitted">
					<not-found 
						v-if="notFound"
						:status="notFoundStatus"
					/>
					<film-result
						v-else
						:title="title"
						:url="url"
						:imgUrl="imgUrl"
					/>
				</div>
			</section>
		</main>
		<goodbyte-footer/>
	</div>
</template>

<script>
	import Logo from './components/Logo.vue';
	import DarkmodeSwitch from './components/DarkmodeSwitch.vue';
	import AboutText from './components/AboutText.vue';
	import SearchBar from './components/SearchBar.vue';
	import AdvancedOptions from './components/AdvancedOptions.vue';
	import LoadingBar from './components/LoadingBar.vue';
	import NotFound from './components/NotFound.vue'
	import FilmResult from './components/FilmResult.vue'
	import GoodbyteFooter from './components/GoodbyteFooter.vue';

	export default 
	{
		name: "App",
		components: 
		{
			Logo,
			DarkmodeSwitch,
			AboutText,
			SearchBar,
			AdvancedOptions,
			LoadingBar,
			NotFound,
			FilmResult,
			GoodbyteFooter
		},
		data()
		{
			return {
				users: '',              // list of users to search
				advancedOptions: {},    // advanced option settings

				info: "",               // json blob from AJAX request
				currentHash: null,      // track most recent request

				loading: false,         // loading statehood
				submitted: true,        // submit statehood

				notFound: false,        // not found error
				emptyintersect: false,  // intersect error
				timeout: false          // timeout error
			}
		},
		created()
		{
			/**
			 * get starting values from url params (if exist)
			 */
			const urlParams = new URLSearchParams(window.location.search);

			const users = urlParams.getAll("u");
			const intersect = urlParams.get("i");
			const ignore = urlParams.get("ignore");

			if (users.length > 0)
			{
				this.users = users.toString();
			}
			
			if (intersect != null)
			{
				this.advancedOptions = { 'selectionMode': 'Intersect' };
			}

			if (ignore != null)
			{
				let ignoreList = ignore.split(',');

				if (ignoreList.includes('unreleased') || ignoreList.includes('true'))
				{
					this.advancedOptions['unreleased'] = false;
				}

				if (ignoreList.includes('shorts'))
				{
					this.advancedOptions['shortFilms'] = false;
				}

				if (ignoreList.includes('feature'))
				{
					this.advancedOptions['featureLength'] = false;
				}
			}

			this.submit();
		},
		methods:
		{
			/**
			 * main function, runs process to get random film
			 */
			submit()
			{
				this.notFound = false;

				if (this.users == "")
				{
					// reset state if form submitted with empty input field
					this.reset();
					return;
				}

				let userlist = this.users.split(/(?:,| )+/); //split input field on space or comma
				// let userlist = inputted.filter(function (el) { return el });
				
				if (userlist.length < 1)
				{
					// second check for non empty input field (only whotespace or commas entered)
					this.submitted = false;
					return;
				}


				document.body.classList.remove("done");
				document.body.classList.add("entered");

				this.submitted = true;
				this.loading = true;

				let ignoreList = [];
				if (this.advancedOptions['unreleased'] == false)
				{
					ignoreList.push('unreleased')
				}
				if (this.advancedOptions['shortFilms'] == false)
				{
					ignoreList.push('shorts')
				}
				if (this.advancedOptions['featureLength'] == false)
				{
					ignoreList.push('feature')
				}


				let apiUrl = "/api?users=" + userlist.join("&users=");
				let clientUrl = "?u=" + userlist.join("&u=");

				if (ignoreList.length > 0)
				{
					apiUrl += "&ignore=" + ignoreList.join(",");
					clientUrl += "&ignore=" + ignoreList.join(",");
				}

				if (this.advancedOptions['selectionMode'] == "Intersect")
				{
					apiUrl += "&intersect=true";
					clientUrl += "&i=true";
				}


				window.history.replaceState(null, null, clientUrl);

				try 
				{
					let vue = this;
					let hash = this.hashCode(apiUrl);
					console.log('url: ' + apiUrl + '\nhash: ' + hash);
					vue.currentHash = hash;
					fetch(apiUrl)
						.then(function (res)
						{
							// check if new request has been sent since submitting
							if (vue.currentHash != hash)
							{
								return "ignoreOldRequest";
							}

							document.body.classList.remove("entered");
							document.body.classList.add("done");

							// bad response
							if (res.status != 200)
							{
								if (res.status == 406)
								{
									vue.emptyintersect = true;
								}

								if (res.status == 502)
								{
									vue.timeout = true;
								}
								
								vue.notFound = true;
								vue.loading = false;

								return "";
							}

							// good response

							setTimeout(function()
							{
								vue.loading = false;
							}, 200); // wait 200ms to load text to allow for image to preload

							return res.json();
						})
						.then(function(json)
						{
							if (json == "ignoreOldRequest")
							{
								return;
							}

							if (!vue.notfound)
							{
								// preload image
								var pre_image = new Image();
								pre_image.src = json.image_url;
							}

							vue.info = json;
						}
					);
				}
				catch (e)
				{
					alert(
						"Something went wrong. Please try again in a moment. Error:" + e,
						"An error occured"
					);
				}
			},

			/**
			 * reset the state
			 */
			reset()
			{
				window.history.replaceState(null, null, "/");
				this.loading = false;
				this.submitted = false;
				document.body.classList.remove("done");
				document.body.classList.remove("entered");
			},

			/**
			 * hashing function to help manage requests
			 */
			hashCode(s)
			{
				let h;
				for (let i = 0; i < s.length; i++)
					h = (Math.imul(31, h) + s.charCodeAt(i)) | 0;
				return h + Date.now();
			}
		},
		computed:
		{
			title: function()
			{
				return this.info.film_name;
			},

			url: function()
			{
				return this.info.slug;
			},

			imgUrl: function()
			{
				return this.info.image_url;
			},

			notFoundStatus: function()
			{
				if (this.emptyintersect)
				{
					return 'no-intersect';
				}
				else if (this.timeout)
				{
					return 'timeout';
				}
				else if (!(this.advancedOptions['unreleased'] && this.advancedOptions['shortFilms'] && this.advancedOptions['featureLength']))
				{
					return 'possibly-ignored';
				}
				else
				{
					return 'other';
				}
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

		transition: color ease-in-out 0.25s, background-color ease-in-out 0.25s;
		margin-top: 30px;
	}

	main 
	{
		min-height: calc(100vh - 100px);
		transform: translateY(40px);
		transition: transform 1.2s cubic-bezier(0.82, 0.01, 0.45, 1);
	}

	.entered main,
	.done main
	{
		transform: none;
	}

	#film-results
	{
		margin: -20px 0;
		transform: translateY(-160px);
		transition: transform 0.3s ease;
	}

	#film-results.advanced-active
	{
		background-color: var(--background);
		position: relative;
		z-index: 99;

		transform: translateY(0px);
		transition: transform 0.3s ease;
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
