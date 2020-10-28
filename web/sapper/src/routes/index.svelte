<script>
	// import Footer from '../components/Footer.svelte';
	import RedIcon from '../components/redIcon.svelte';

	let header = "Welcome";
	let osTechList = ["GO", "go-Micro", "Svelte", "NATS"];
	let dbTechList = ["ArangoDB", "PostgreSQL", "Redis", "TimescaleDB"];
	let description = "This project provides a sample full stack microservices implementation. goTemp was built by Juan Peredo at ";
	let bolbeck = {label:"Bobleck.com", address: "https://www.bolbeck.com"}
	let repo = "https://bitbucket.org/Bolbeck/gotemp"

</script>

<style>
	h1, p {
		margin: 0 auto;
	}

	h1 {
		font-size: 2.8em;
		font-weight: 500;
		margin: 0 0 0.5em 0;
	}

	img {
		width: 100%;
		margin: 3em 0 3em 0;
	}

	.pmar {
		margin: 1em auto;
	}

	@media (min-width: 480px) {
		h1 {
			font-size: 2em;
		}
	}

	section {
		padding-top: 5rem;
		padding-bottom: 5rem;
	}

</style>

<svelte:head>
	<title>goTemp</title>
</svelte:head>

<section style="background-image: url('mainbackground.jpg'); background-repeat: no-repeat; background-position: center center; background-size: cover">
	<div class="container" >
		<div class="row" >
			<div class="col-lg-6 order-1 order-lg-1">
				<img src="gotempwelcome.png" alt="Welcome to goTemp" class="img-fluid">
			</div>
			<div class="col-lg-6 order-2 order-lg-2">
				<h1>{header}</h1>
				<p class="lead pmar">{description}
					<a href="{bolbeck.address}">
						{bolbeck.label}
					</a>
				</p>
				<p class="pmar">Try exploring the application using the menu items at the top of the page. If you are curious about the tech stack, feel free to review the <a href="https://bitbucket.org/Bolbeck/gotemp">repo documentation in bitbucket </a> for more info. <br/>
					If you find any issues or problems please create an issue or a pull request.
				</p>
				<p>
					<a href="{repo}" class="btn btn-secondary mr-2">Repo</a>
				</p>
			</div>
		</div>
	</div>
</section>

<section>
	<div class="container">
		<h2>Technology stack</h2>
		<p class="text-muted mb-5">Below is the list of technologies used to build this demo.</p>

		<div class="row">
			<div class="col-sm-6 col-lg-4 mb-3">
				<RedIcon  iconName="fa-desktop" />
				<h5>Frontend</h5>
				<p class="text-muted">Language: Javascript</p>
				<p class="text-muted">Framework: Svelte/Sapper</p>
			</div>
			<div class="col-sm-6 col-lg-4 mb-3">
				<RedIcon  iconName="fa-server" />
				<h5>Backend</h5>
				<p class="text-muted">Language: Go</p>
				<p class="text-muted">Framework: go-Micro</p>
			</div>
			<div class="col-sm-6 col-lg-4 mb-3">
				<RedIcon  iconName="fa-database" />
				<h5>Databases</h5>
				<p class="text-muted">PostgreSql: Relational data storage</p>
				<p class="text-muted">ArangoDB: Multi-model for master data</p>
				<p class="text-muted">Timescale: Time-series for audit data</p>
				<p class="text-muted">Redis: Key-value DB for caching</p>
				<br>
			</div>
		</div>
		<div class="row">
			<div class="col-sm-6 col-lg-4 mb-3">
				<RedIcon  iconName="fa-cubes" />
				<h5>Orchestration</h5>
				<p class="text-muted">Docker</p>
				<p class="text-muted">Kubernetes running in Minikube</p>
			</div>
			<div class="col-sm-6 col-lg-4 mb-3">
				<RedIcon  iconName="fa-comments" />
				<h5>Communication</h5>
				<p class="text-muted">gRPC: Inter-service communication</p>
				<p class="text-muted">JSON: External communication</p>
			</div>
			<div class="col-sm-6 col-lg-4 mb-3">
				<RedIcon  iconName="fa-book" />
				<h5>Other</h5>
				<p class="text-muted">NATS: Pub/sub</p>
				<p class="text-muted">NGINX: K8s ingress controller</p>
			</div>
		</div>
	</div>
</section>


<!--<Footer />-->