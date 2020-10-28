<script>
	// import Footer from '../components/Footer.svelte'
	import IconWtext from '../components/iconWtext.svelte'
	import Col from 'sveltestrap/src/Col.svelte'
	import Row from 'sveltestrap/src/Row.svelte'

	let header = "Welcome";
	let osTechList = ["GO", "go-Micro", "Svelte", "NATS"];
	let dbTechList = ["ArangoDB", "PostgreSQL", "Redis", "TimescaleDB"];
	let description = "This project provides a sample full stack microservices implementation. goTemp was built by Juan Peredo at ";
	let bolbeck = {label:"Bobleck.com", address: "https://www.bolbeck.com"}
	let repo = "https://bitbucket.org/Bolbeck/gotemp"
	let frontend = {header: "Frontend", iconName: "fa-desktop", techList: ["Language: Javascript","Framework: Svelte/Sapper"]}
	let backend = {header: "Backend", iconName: "fa-server", techList: ["Language: Go","Framework: go-Micro"]}
	let database = {header: "Databases",
					iconName: "fa-database",
					techList: ["PostgreSql: Relational data storage",
						"ArangoDB: Multi-model for master data",
						"Timescale: Time-series for audit data",
						"Redis: Key-value DB for caching"
					]}
	let orchestration = {header: "Orchestration", iconName: "fa-cubes", techList: ["Docker","Kubernetes running in Minikube"]}
	let communication = {header: "Communication", iconName: "fa-comments", techList: ["gRPC: Inter-service communication","JSON: External communication"]}
	let other = {header: "Other", iconName: "fa-book", techList: ["NATS: Pub/sub","NGINX: K8s ingress controller"]}


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

	.sectionMain {
		background-repeat: no-repeat;
		background-position: center center;
		background-size: cover;
	}
</style>

<svelte:head>
	<title>goTemp</title>
</svelte:head>

<section class="sectionMain" style="background-image: url('mainbackground.jpg');">
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
				<p class="pmar">Try exploring the application using the menu items at the top of the page. If you are curious about the tech stack, feel free to review the repo documentation in bitbucket for more info. <br/>
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
			<IconWtext data="{frontend}" />
			<IconWtext data="{backend}" />
			<IconWtext data="{database}" />
		</div>
		<div class="row">
			<IconWtext data="{orchestration}" />
			<IconWtext data="{communication}" />
			<IconWtext data="{other}" />
		</div>
	</div>
</section>


<!--<Footer />-->