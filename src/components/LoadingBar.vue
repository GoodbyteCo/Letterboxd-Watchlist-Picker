<template>
	<div>
		<h2>Loading Film</h2>
		<div id="loading-bar"></div>
		<p>
			Is there a movie you're secretly rooting for? 
			Choose that one! I give you permission.
		</p>
	</div>
</template>

<script>
export default
{
	name: 'LoadingBar'
}
</script>

<style scoped>
	h2
	{
		margin-top: 5rem;
		font-size: 1.2rem;
	}

	p
	{
		max-width: 40ch;
		margin: 0rem auto 1rem;
		padding: 0 2rem;
	}

	#loading-bar
	{
		background: var(--off-white);
		border-radius: 10px;

		margin: 1rem auto;
		max-width: 90%;
		width: 600px;
		height: 3px;
	}

	#loading-bar::after
	{
		content: "";
		display: block;
		height: 3px;
		width: 100%;
		border-radius: 4px;

		background: var(--primary);
		
		animation-fill-mode: forwards;
		transform: scaleX(0);
		transform-origin: left;
		transition: transform 1.4s ease;
	}

	.entered #loading-bar::after
	{
		animation: load 1.4s ease;
		animation-fill-mode: forwards;
	}

	.done #loading-bar::after
	{
		transform: scaleX(0.8);
		animation: load-finish 0.2s ease-in;
		animation-fill-mode: forwards;
		transition: transform 0.2s ease-in;
	}

	@keyframes load
	{
		0% { transform: scaleX(0); }
		100% { transform: scaleX(0.8); }
	}

	@keyframes load-finish
	{
		100% { transform: scaleX(1); }
	}
</style>
