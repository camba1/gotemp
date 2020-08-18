<script>
	import Navbar from '../components/Navbar.svelte';

	export let segment;
</script>

<style>
	main {
		position: relative;
		max-width: 78em;
		background-color: white;
		padding: 1em;
		margin: 0 auto;
		box-sizing: border-box;
		/*overflow:scroll;*/
	}
	/*main {*/
	/*	text-align: center;*/
	/*	padding: 1em;*/
	/*	max-width: 240px;*/
	/*	margin: 0 auto;*/
	/*}*/


</style>

<Navbar {segment}/><br/><br/>

<main>

	<slot></slot>

</main>
