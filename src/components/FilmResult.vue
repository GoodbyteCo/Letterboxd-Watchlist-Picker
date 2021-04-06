
<template>
	<div class="result">
		<a
			v-bind:href="url"
			:style="{ backgroundImage: 'url(' + imgUrl + ')' }"
			class="film-cover"
			alt="film poster"
			v-focus-visible>
		</a>
		<div>
			<p>
				You should watch
			</p>
			<a 
				v-bind:href="url" 
				class="title"
				v-focus-visible
			>
				{{ title }}
			</a>
		</div>
	</div>
</template>

<script>
	export default
	{
		name: 'FilmResult',
		props:
		[
			'title',
			'url',
			'imgUrl'
		]
	}
</script>

<style scoped>
	.result
	{
		display: grid;
		grid-template-rows: 345px auto;
		gap: 20px;
		text-align: center;
		max-width: 400px;
		width: 90%;
		margin: 60px auto;
		margin-top: 30px;
	}

	.film-cover
	{
		display: block;
		width: 230px;
		height: 345px;
		border-radius: 4px;
		margin: auto;
		outline: none;

		background-color: var(--off-white);
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.35), 0 0 2px 1px rgb(0 0 0 / 5%);
		transition: background-color ease-in-out 0.25s;
		
	}

	.dark .film-cover
	{
		background-color: var(--tertiary);
	}

	.film-cover:hover,
	.film-cover[v-focus-visible=true]:focus
	{
		box-shadow: inset 0 0 0 3px var(--primary),
			0 1px 3px rgba(0, 0, 0, 0.35),
			0 0 2px 1px rgb(0 0 0 / 5%);
	}

	p
	{
		max-width: 60ch;
		padding: 0 2rem;
		margin: 1rem auto 0.5rem;
	}

	.title
	{
		color: var(--foreground);
		font-size: 1.5rem;
		font-weight: bold;
		margin: 0;
	}

	.title:hover,
	.title[v-focus-visible=true]:focus
	{
		color: var(--primary);
	}
</style>
