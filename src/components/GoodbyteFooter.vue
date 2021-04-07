<template>
	<footer>
		<a data-goatcounter-click="support-us" data-goatcounter-title="Support Us" href="https://ko-fi.com/goodbyte" class="support">
			Support this site
		</a>
		<p>
			<a href="http://goodbyte.ca/">Goodbyte</a> &copy; 2020–{{ year }}
		</p>
	</footer>
</template>

<script>
	export default
	{
		name: 'GoodbyteFooter',
		computed:
		{
			year()
			{
				return new Date().getFullYear();
			}
		}
	}
</script>

<style scoped>
	footer
	{
		text-align: center;
		margin: 0 auto;
		transform: translateY(-20px);
		
		font-size: 12px;
		letter-spacing: 0.03em;
	}

	a.support
	{
		display: inline-block;
		padding: 0.3em 0.5em;
		box-sizing: border-box;
		overflow: hidden;
		height: 2.4em;

		border: 3px solid var(--tertiary);
		border-radius: 4px;

		color: var(--primary);
		text-align: center;
		text-transform: uppercase;
		font-weight: 900;
		text-decoration: none;
		white-space: nowrap;

		width: 0px;
		opacity: 0;
		margin-right: 0px;
		pointer-events: none;

		transition: opacity 0.8s ease,
			width 0.8s ease-in,
			margin 0.8s ease-in;
	}
	
	.done a.support,
	[v-focus-visible=true] a.support:focus
	{
		width: 150px;
		opacity: 1;
		margin-right: 28px;
		pointer-events: all;

		transition: opacity 0.6s ease,
			width 0.6s ease-out,
			margin 0.6s ease-out;
	}


	a.support:hover,
	[v-focus-visible=true] a.support:focus
	{
		background: var(--primary);
		border-color: var(--primary);
		color: var(--background);
	}

	p
	{
		margin: 0;
		display: inline-block;
		opacity: 0.6;

		transform: translateX(-18px) translateY(-10px);
	}

	p a
	{
		color: inherit;
		text-decoration: none;
		border-radius: 2px;

		background-size: 100% 1px;
		background-position: 0 1.3em;
		background-repeat: no-repeat;
	}

	p a:hover,
	[v-focus-visible=true] p a:focus
	{
		background-image: linear-gradient(var(--foreground), var(--foreground));
	}

	[v-focus-visible=true] a:focus,
	[v-focus-visible=true] a.support:focus
	{
		opacity: 1;
		box-shadow: 0 0 0 3px var(--background), 0 0 0 5px var(--primary);
		outline: none;
	}
</style>
