<script>
	import Footer from '../components/Footer.svelte';
</script>

<style>
	h1, figure, p {
		text-align: center;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		font-size: 2.8em;
		text-transform: uppercase;
		font-weight: 700;
		margin: 0 0 0.5em 0;
	}

	figure {
		margin: 0 0 1em 0;
	}

	img {
		width: 100%;
		max-width: 400px;
		margin: 0 0 1em 0;
	}

	p {
		margin: 1em auto;
	}

	@media (min-width: 480px) {
		h1 {
			font-size: 4em;
		}
	}

</style>

<svelte:head>
	<title>goTemp</title>
</svelte:head>

<h1>Welcome to goTemp</h1>

<figure>
	<img alt='Success Kid' src='successkid.jpg'>
	<figcaption>Now, let's have fun with some micro-services!</figcaption>
</figure>

<p><strong>Try exploring some of the menu items and review the bitbucket docs for more info. <br/>
			If you have any issues please create an issue or a pull request.</strong></p>

<Footer />