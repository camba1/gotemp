<script>
    import { Col, Container, Row } from "sveltestrap"

    let header = "goTemp";
    let osTechList = ["GO", "go-Micro", "Svelte", "NATS"];
    let dbTechList = ["ArangoDB", "PostgreSQL", "Redis", "TimescaleDB"];
    let description = "This project provides a sample microservices implementation. goTemp was built by Juan Peredo at ";

</script>

<style>

    footer {
        position: absolute;
        left: 0;
        bottom: 0;
        height: 150px;
        width: 100%;
        overflow: hidden;
    }

</style>

<footer class="p-3 my-3 bg-light">

    <Container>
        <Row>
            <Col></Col>
            <Col class="col-3">
                <h6>{header}</h6>
                <p>
                    {description}<a href="https://www.bolbeck.com">Bobleck.com</a>
                </p>
            </Col>
            <Col class="col-3">
                <h6>Powered by open source</h6>
                <ul class="list-unstyled">
                        {#each osTechList as list}
                            <li>{list}</li>
                        {/each}
                </ul>
            </Col>
            <Col class="col-3">
                <h6>Powered by open source DBs</h6>
                <ul class="list-unstyled">
                    {#each dbTechList as list}
                        <li>{list}</li>
                    {/each}
                </ul>
            </Col>
            <Col></Col>
        </Row>
    </Container>
</footer>



