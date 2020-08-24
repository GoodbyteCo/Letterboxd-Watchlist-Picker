<template>
	<footer>
		<a
			data-goatcounter-click="support-us"
			data-goatcounter-title="Support US"
			href="https://ko-fi.com/goodbyte"
			class="support"
		>Support this site</a>
		<a href="http://goodbyte.ca/">Goodbyte</a> &copy; 2020
	</footer>
</template>

<script>
export default {
	name: "Footer",
};
</script>

<style scoped>
footer {
	text-align: center;
	margin: 0 auto 2rem;
	opacity: 0.6;
	font-size: 12px;
	letter-spacing: 0.03em;
}

a {
	color: inherit;
	text-decoration: none;
}

a:hover {
	text-decoration: underline;
}

a.support {
	display: block;
    padding: 0.5rem;
    text-align: center;
    margin: 0.5rem auto;
    width: 120px;
    border-radius: 4px;
    text-transform: uppercase;
    font-weight: 900;
    text-decoration: none;
    color: #00bdfa;
    border: 3px solid #00bdfa;
	opacity: 0;
	pointer-events: none;
	transition: opacity 0.6s ease 0.2s;
}

a.support:hover {
	background: #00bdfa;
	color: white;
}

.done a.support {
	opacity: 1;
	pointer-events: all;
}
</style>
