<script>
	import Navbar from "./Components/Navbar.svelte";
	import Footer from "./Components/footer.svelte";
	import Login from "./Components/login.svelte";
	import Products from "./Components/product/products.svelte";
	export let name;

	function handleMessage(event) {
		alert(event.detail.text);
	}

</script>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		/*text-transform: uppercase;*/
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>

<Navbar/>

<main>
	<h1>{name}</h1>
	<p>A sample microservices open source project developed at <a href="https://bolbeck.com">Bolbeck</a> to showcase some seriously cool tech.</p>
</main>

<Login on:message={handleMessage}/>

<Products/>


<!--<Footer/>-->