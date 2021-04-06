<template>
	<div>
			>
					/>
				<div
					id="advanced-section"
					aria-haspopup="true"
					aria-expanded="false"
				>
					<button v-on:click="activateAdvanced()" id="tertiary">
						<span>Advanced Search</span>
					</button>
					<div id="advanced">
						<input
							type="radio"
							id="union"
							value="Union"
							name="union-intersect"
							v-model="selectionMode"
							tabindex="-1"
							selected
						/>
						<label for="union">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								viewBox="0 0 131 87"
							>
								<g stroke="#000" stroke-width="3">
									<g transform="translate(-796 -208)">
										<circle
											cx="839.7"
											cy="251.7"
											r="43.7"
											stroke="none"
										/>
										<circle
											cx="839.7"
											cy="251.7"
											r="42.2"
											fill="none"
										/>
									</g>
									<g transform="translate(-796 -208)">
										<circle
											cx="883.3"
											cy="251.7"
											r="43.7"
											stroke="none"
										/>
										<circle
											cx="883.3"
											cy="251.7"
											r="42.2"
											fill="none"
										/>
									</g>
									<g
										fill="none"
										transform="translate(-796 -208)"
									>
										<circle
											cx="839.7"
											cy="251.7"
											r="43.7"
											stroke="none"
										/>
										<circle
											cx="839.7"
											cy="251.7"
											r="42.2"
										/>
									</g>
									<path
										fill="#fff"
										d="M15 13l73 73zM2 40l46 46zm35 45L2 50zM20 8l77 77zm56 76L10 18zM3 31l53 53zM27 5l77 77zM6 24l57 57zm16 56L8 66zM35 3l76 76zm9-1l72 72zm11 1l66 66zm13 3l57 57zm7-3l53 53zm9-1l45 45zm10 0l35 35zm16 6l13 13z"
									/>
								</g>
							</svg>
							Union
						</label>
						<input
							type="radio"
							id="intersect"
							value="Intersect"
							name="union-intersect"
							tabindex="-1"
							v-model="selectionMode"
						/>
						<label for="intersect">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								viewBox="0 0 131 87"
							>
								<g stroke="#000" stroke-width="3">
									<g transform="translate(-786 -52)">
										<circle
											cx="829.7"
											cy="95.7"
											r="43.7"
											stroke="none"
										></circle>
										<circle
											cx="829.7"
											cy="95.7"
											r="42.2"
											fill="none"
										></circle>
									</g>
									<g transform="translate(-786 -52)">
										<circle
											cx="873.3"
											cy="95.7"
											r="43.7"
											stroke="none"
										></circle>
										<circle
											cx="873.3"
											cy="95.7"
											r="42.2"
											fill="none"
										></circle>
									</g>
									<g
										fill="none"
										transform="translate(-786 -52)"
									>
										<circle
											cx="829.7"
											cy="95.7"
											r="43.7"
											stroke="none"
										></circle>
										<circle
											cx="829.7"
											cy="95.7"
											r="42.2"
										></circle>
									</g>
									<path
										fill="#fff"
										d="M47 55l22 22zm-2-12l29 29zm1-9l33 33zm3-7l33 33zm3-7l33 33zm5-5l29 29zm5-5l22 22z"
									></path>
								</g>
							</svg>
							Intersect</label
						>
						<div id="ignore-box">
							<input
								type="checkbox"
								id="ignore"
								v-model="ignoreChecked"
							/>
							<label for="ignore">
								Ignore unreleased films
								<span>
									Removes all films released this year or in
									the future from results.
								</span>
							</label>
						</div>
					</div>
				</div>
			</div>
			<div id="advanced-show">
				<div v-if="loading">
					<h2>Loading Film</h2>
					<div id="loadbar"></div>
					<p>
						Is there a movie you're secretly rooting for? Choose
						that one! I give you permission.
					</p>
				</div>
				<div v-else-if="submitted">
					<div v-if="notfound">
						<h2>Nothing Found</h2>
						<p v-if="emptyintersect">
							The intersection between those two lists is empty.
						</p>
						<p v-else-if="ignoreChecked">
							There were no films found in that list. It may be
							empty, private, or only contain films
							not-yet-released (films released in the current year
							are also excluded).
						</p>
						<p v-else-if="timeout">
							Sorry your list was too powerful and we timed out.
							Try another list
						</p>
						<p v-else>
							Sorry, that watchlist is empty, private, or doesn't
							exist at all.
						</p>
						<img
							id="poe"
							width="250"
							src="https://watchlistpicker.com/poe.gif"
							alt="from the movie Kung-Fu Panda, protaganist Poe looks down at an empty scroll."
						/>
					</div>
					<div v-else id="container">
						<a
							v-bind:href="url"
							:style="{ backgroundImage: 'url(' + img_url + ')' }"
							class="film-cover"
							alt="film poster"
						></a>
						<div>
							<p class="you-should">You should watch</p>
							<a
								id="title-link"
								v-bind:href="url"
								class="title"
								>{{ name }}</a
							>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: "Film",
	data() {
		return {
			users: "", //sting of requested users binding for input box
			info: "", //json blob fotten from AJAX request
			notfound: false, //Boolean for when user is not found or a watchlist is empty
			emptyintersect: false,
			loading: false, //Boolean fot loading state
			submitted: false, //Boolean for if the form has been submitted
			selectionMode: "Union",
			advancedOpen: false,
			currentHash: null,
			timeout: false
		};
	},
	//Query to see if users has passed url params to make quick request
	created() {
		const queryString = window.location.search;
		console.log(queryString);
		const urlParams = new URLSearchParams(queryString);
		const users = urlParams.getAll("u");
		const inter = urlParams.get("i");
		const ignore = urlParams.get("ignore");
		if (users.length > 0) {
			this.users = users.toString();
		}
		if (inter != null) {
			this.selectionMode = "Intersect";
		}
		if (ignore != null) {
			this.ignoreChecked = true;
		}
		this.submit();
	},
	methods: {
		//Main function to make request for random film
		submit() {
			this.notfound = false;
			if (this.users == "") {
				//Reset state and if form submitted with empty input field
				window.history.replaceState(null, null, "/");
				this.loading = false;
				this.submitted = false;
				document.body.classList.remove("done");
				document.body.classList.remove("entered");
				return;
			}
			let inputted = this.users.split(/(?:,| )+/); //split input field on space or comma
			let userlist = inputted.filter(function(el) {
				return el;
			});
			if (userlist.length < 1) {
				// second check for non empty input field probably not required
				this.submitted = false;
				return;
			}
			document.body.classList.remove("done");
			document.body.classList.add("entered");
			this.submitted = true;
			this.loading = true;
			//TODO add window handeling state for ignore current year
			if (this.selectionMode == "Intersect" && this.ignoreChecked) {
				window.history.replaceState(
					null,
					null,
					"?u=" + userlist.join("&u=") + "&i=true&ignore=true"
				); //add url param for users being queryied for discoverbilty of this feature
			} else if (this.selectionMode == "Intersect") {
				window.history.replaceState(
					null,
					null,
					"?u=" + userlist.join("&u=") + "&i=true"
				);
			} else if (this.ignoreChecked) {
				window.history.replaceState(
					null,
					null,
					"?u=" + userlist.join("&u=") + "&ignore=true"
				);
			} else {
				window.history.replaceState(
					null,
					null,
					"?u=" + userlist.join("&u=")
				); //add url param for users being queryied for discoverbilty of this feature
			}
			console.log(userlist);

			//Generate proper url for request

			let url = "/api?users=" + userlist.join("&users=");
			console.log(this.selectionMode);
			if (this.selectionMode == "Intersect") {
				url += "&intersect=true";
			}
			if (this.ignoreChecked) {
				url += "&ignore_unreleased=true";
			}
			try {
				let vue = this;
				console.log(url);
				let hash = this.hashCode(url);
				console.log(hash);
				vue.currentHash = hash;
				fetch(url)
					.then(function(res) {
						if (vue.currentHash != hash) {
							return "ignoreOldRequest";
						}
						document.body.classList.remove("entered");
						document.body.classList.add("done");

						if (res.status != 200) {
							if (res.status == 406) {
								vue.emptyintersect = true;
							}
							if (res.status == 502) {
								vue.timeout = true;
							}
							//if request fails set state to failed stated
							vue.notfound = true;
							vue.loading = false;
							return "";
						}

						//wait 200ms to load text to allow for image to preload
						setTimeout(function() {
							vue.loading = false;
						}, 200);

						return res.json();
					})
					.then(function(json) {
						if (json == "ignoreOldRequest") {
							return;
						}
						if (!vue.notfound) {
							// Preload image
							var pre_image = new Image();
							pre_image.src = json.image_url;
						}
						vue.info = json;
					});
			} catch (e) {
				this.$alert(
					"Something went wrong. Please try again in a moment. Error:" +
						e,
					"An error occured"
				);
			}
		},
		activateAdvanced() {
			var large = document.getElementById("advanced-show");
			var element = document.getElementById("advanced");
			large.classList.toggle("advanceactive");
			element.classList.toggle("active");
			if (this.advancedOpen) {
				document.getElementById("union").tabIndex = -1;
				document.getElementById("intersect").tabIndex = -1;
				document
					.getElementById("advanced-section")
					.setAttribute("ariaexpanded", "false");
				this.advancedOpen = false;
			} else {
				document.getElementById("union").tabIndex = 0;
				document.getElementById("intersect").tabIndex = 0;
				document
					.getElementById("advanced-section")
					.setAttribute("ariaexpanded", "true");
				this.advancedOpen = true;
			}
		},
		hashCode(s) {
			let h;
			for (let i = 0; i < s.length; i++)
				h = (Math.imul(31, h) + s.charCodeAt(i)) | 0;
			return h + Date.now();
		}
	},
	computed: {
		filmNotFound: function() {
			return this.notfound;
		},
		url: function() {
			return this.info.slug;
		},
		img_url: function() {
			return this.info.image_url;
		},
		name: function() {
			return this.info.film_name;
		}
	}
};
</script>

<style scoped>
#title-link {
	font-size: 1.5rem;
	color: #415569;
	transition: color ease-in-out 0.25s;
}

a.title {
	font-weight: bold;
	font-size: 2rem;
	margin: 0;
}

#title-link:hover,
#title-link:focus {
	color: #40bcf4;
}

#poe {
	margin: 30px 0;
	border-radius: 4px;
}

a {
	color: #415569;
	transition: color ease-in-out 0.25s;
}

a:hover,
a:focus {
	color: #40bcf4;
	outline: none;
}

#advanced {
	opacity: 0;
	transition: 0.3s ease;
	transform: translateY(0px);
}

#advanced label {
	display: inline;
	visibility: visible;
}

#advanced span {
	opacity: 0.6;
	display: block;
}

#ignore-box {
	padding: 2rem 0.8rem;
}

#ignore-box label {
	max-width: 200px;
	text-align: left;
	display: inline-block;
	margin: 10px;
}

#ignore-box span {
	font-size: 12px;
}

#ignore {
	transform: translateY(-2rem);
	display: inline-block;
}

.active#advanced {
	opacity: 1;
	transition: 0.3s ease;
	transform: translateY(5px);
}

#advanced-show {
	transform: translateY(-180px);
	transition: transform 0.3s ease;
}

#advanced-show.advanceactive {
	transform: translateY(-20px);
	transition: transform 0.3s ease;
}

#advanced svg {
	width: 40px;
	transform: translateY(8.5px);
	margin-right: 4px;
	fill: transparent;
}

#advanced label {
}

#advanced > label {
	display: inline;
	visibility: visible;
	background: rgb(64 188 244 / 0.5);
	transition: background ease-in-out 0.25s;
	padding: 10px;
	border-radius: 4px;
	font-weight: bold;
	color: rgb(0 0 0 / 70%);
	border: 2px solid transparent;
	cursor: pointer;
	margin: 4px;
}

.dark #advanced label {
	background: #526e89;
	transition: background ease-in-out 0.25s;
}

.dark #ignore-box > label {
	background: unset;
}

#advanced input[type="radio"]:checked + label {
	background: #40bcf4;
	transition: border-color ease-in-out 0.25s;
	border-color: #000;
	color: #fff;
}

.dark #advanced input[type="radio"]:checked + label {
	border-color: #fff;
	transition: border-color ease-in-out 0.25s;
}

#advanced input[type="radio"]:checked + label svg {
	fill: #fff;
}

#advanced input[type="radio"]:focus + label {
	box-shadow: 0 0 0 3px #fff, 0 0 0 5px #1caff2;
}

.dark #advanced input[type="radio"]:focus + label {
	box-shadow: 0 0 0 3px #14181d, 0 0 0 5px #1caff2;
}

input#union,
input#intersect {
	opacity: 0;
	width: 0px;
	margin: 0;
}

h1 {
	font-size: 1.5rem;
	transform: scale(2) translateY(15px);
	transition: transform 1.2s cubic-bezier(0.82, 0.01, 0.45, 1);
}

#logo {
	transform: scale(1.5) translateY(15px);
	transition: transform 1.2s cubic-bezier(0.82, 0.01, 0.45, 1);
}

a#logo:focus-visible {
	box-shadow: 0 0 0 2px #fff, 0 0 0 3.5px #1caff2;
}

.dark a#logo:focus-visible {
	box-shadow: 0 0 0 2px #14181d, 0 0 0 3.5px #1caff2;
}

#app a:focus-visible {
	box-shadow: 0 0 0 3px #fff, 0 0 0 5px #1caff2;
	border-radius: 3px;
	outline: none;
}

.dark #app a:focus-visible {
	box-shadow: 0 0 0 3px #14181d, 0 0 0 5px #1caff2;
}

.entered h1,
.entered #logo,
.entered .hello,
.done h1,
.done #logo,
.done .hello {
	transform: none;
}

h2 {
	margin-top: 5rem;
	font-size: 1.2rem;
}

h3 {
	margin: 40px 0 0;
}

::placeholder {
	opacity: 0.6;
}

#loadbar {
	margin: 1rem auto;
	max-width: 90%;
	width: 600px;
	height: 3px;
	background: #ebebeb;
	border-radius: 10px;
}

#loadbar::after {
	content: "";
	display: block;
	background: #40bcf4;
	height: 3px;
	width: 100%;
	border-radius: 4px;
	animation-fill-mode: forwards;
	transform: scaleX(0);
	transform-origin: left;
	transition: transform 1.4s ease;
}

.entered #loadbar::after {
	animation: load 1.4s ease;
	animation-fill-mode: forwards;
}

.done #loadbar::after {
	transform: scaleX(0.8);
	animation: load-finish 0.2s ease-in;
	animation-fill-mode: forwards;
	transition: transform 0.2s ease-in;
}

@keyframes load {
	0% {
		transform: scaleX(0);
	}
	100% {
		transform: scaleX(0.8);
	}
}

@keyframes load-finish {
	100% {
		transform: scaleX(1);
	}
}

.film-cover {
	margin: auto;
	box-shadow: 0 1px 3px rgba(0, 0, 0, 0.35), 0 0 2px 1px rgb(0 0 0 / 5%);
	background-color: #ebebeb;
	transition: background-color ease-in-out 0.25s;
	display: block;
	width: 230px;
	height: 345px;
	border-radius: 4px;
}

.film-cover:hover,
.film-cover:focus {
	box-shadow: inset 0 0 0 3px #40bcf4, 0 1px 3px rgba(0, 0, 0, 0.35),
		0 0 2px 1px rgb(0 0 0 / 5%);
	outline: none;
}

#container {
	display: grid;
	grid-template-rows: 345px auto;
	gap: 20px;
	text-align: center;
	max-width: 400px;
	width: 90%;
	margin: 60px auto;
	margin-top: 30px;
}

button {
	font-family: Avenir, Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	padding: 0 1rem;
	line-height: 2.8rem;
	border: 0;
	background: black;
	color: white;
	text-transform: uppercase;
	cursor: pointer;
	border-radius: 0 4px 4px 0;
	outline: none;
	font-size: 0.8rem;
	transform: translateY(-0.12rem);
	display: inline-block;
	font-weight: 900;
	letter-spacing: 0.04em;
}

.dark button {
	background: #526e89;
}

button:hover,
button:focus,
button:focus-within {
	background: #40bcf4;
}

#tertiary {
	background: none;
	color: #415569;
	transition: color ease-in-out 0.25s;
	display: block;
	margin: auto;
}

.dark #tertiary {
	color: #526e89;
	transition: color ease-in-out 0.25s;
}

#tertiary:hover,
#tertiary:focus {
	text-decoration: underline;
}

#tertiary span {
	text-transform: none;
	font-weight: 400;
}

#tertiary:focus-visible span {
	text-decoration: underline;
	box-shadow: 0 0 0 2px #fff, 0 0 0 3.5px #1caff2;
	border-radius: 3px;
	padding: 3px 8px;
}

.dark #tertiary:focus-visible span {
	box-shadow: 0 0 0 2px #14181d, 0 0 0 3.5px #1caff2;
}

/* #advanced-section {
	margin: auto;
	text-align: left;
} */

.userfield {
	font-family: Avenir, Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	min-width: 220px;
	border: 0;
	background: #ebebeb;
	padding: 0 1rem;
	line-height: 2.8rem;
	font-size: 1.1rem;
	outline: none;
	border-radius: 4px 0 0 4px;
}

.userfield:active,
.userfield:focus,
.userfield:focus-within {
	box-shadow: inset 0 0 0 3px #40bcf4;
}

@media screen and (max-width: 360px) {
	.userfield {
		min-width: 0px;
		width: 50%;
	}
}

label {
	visibility: hidden;
	display: block;
}

p {
	max-width: 60ch;
	margin: 0rem auto 1rem;
	padding: 0 2rem;
}

p.you-should {
	margin: 1rem auto 0.5rem;
}

/* logo stuff */
#logo {
	display: block;
	text-decoration: none;
	width: 78px;
	height: 25px;
	overflow: hidden;
	border: 2px solid black;
	transition: border ease-in-out 0.25s;
	border-radius: 4px;
	margin: auto;
}

.entered #logo svg {
	animation: spin 0.2s linear infinite;
}

#logo svg {
	transform: translateY(-16.3px);
	animation: finish 0.8s cubic-bezier(0, 0.7, 0.39, 1.35);
}

#logo svg:first-child {
	animation-delay: 0.05s;
	transform: translateY(-16.3px);
}

#logo svg:last-child {
	animation-delay: 0.1s;
	transform: translateY(-16.3px);
}

/* @media (prefers-color-scheme: dark) { */

.dark #logo {
	border: 2px solid white;
	transition: border ease-in-out 0.25s;
}

.dark a {
	color: #76a0ca;
	transition: color ease-in-out 0.25s;
}

.dark #title-link {
	color: #76a0ca;
	transition: color ease-in-out 0.25s;
}

.dark .film-cover {
	background-color: #76a0ca;
	transition: background-color ease-in-out 0.25s;
}

/* } */

@keyframes spin {
	0% {
		transform: translateY(-103px);
	}
	100% {
		transform: translateY(0px);
	}
}

@keyframes finish {
	0% {
		transform: translateY(-100px);
	}
	100% {
		transform: translateY(-16.3px);
	}
}
</style>
