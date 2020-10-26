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
		/*text-transform: uppercase;*/
		font-weight: 700;
		margin: 0 0 0.5em 0;
	}

	figure {
		margin: 0 0 1em 0;
	}

	img {
		width: 100%;
		max-width: 400px;
		margin: 3em 0 3em 0;
	}

	p {
		margin: 1em auto;
	}

	@media (min-width: 480px) {
		h1 {
			font-size: 3em;
		}
	}

</style>

<svelte:head>
	<title>goTemp</title>
</svelte:head>

<h1>Welcome to </h1>

<figure>
	<img alt='Welcome to goTemp' src='gotempwelcome.png'>
	<figcaption>Now, let's have fun with some micro-services!</figcaption>
</figure>

<p>Try exploring some of the menu items above and review the <a href="https://bitbucket.org/Bolbeck/gotemp">bitbucket docs</a> for more info. <br/>
			If you have any issues please create an issue or a pull request.
</p>

<Footer />