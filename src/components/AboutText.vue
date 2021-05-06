<template>
	<p>
		Enter your <a href="https://letterboxd.com/">Letterboxd</a> username to get a random film off of your watchlist. Add multiple usernames by separating with a space, and enter public lists in "username/list-title" format. See <a href="https://github.com/GoodbyteCo/Letterboxd-Watchlist-Picker/blob/main/README.md#frequently-asked-questions" v-focus-visible>FAQs</a> for more.
	</p>
</template>

<script>
	export default
	{
		name: 'AboutText'
	}
</script>

<style scoped>
	a
	{
		color: var(--foreground);

		text-decoration: none;
		background-image: linear-gradient(var(--foreground), var(--foreground));
		background-size: 100% 1px;
		background-position: 0 1.2em;
		background-repeat: no-repeat;

		transition: color ease-in-out 0.25s; /* darkmode transition */
	}

	a:hover,
	[v-focus-visible=true] a:focus
	{
		color: var(--primary);
		background-image: linear-gradient(var(--primary), var(--primary));
		transition: none;
		outline: none;
	}

	[v-focus-visible=true] a:focus
	{
		box-shadow: 0 0 0 3px var(--background), 0 0 0 5px var(--primary);
		border-radius: 3px;
	}

	p
	{
		max-width: 60ch;
		margin: 0rem auto 1rem;
		padding: 1.5rem 2rem 0;
	}
</style>
