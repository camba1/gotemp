<script>
    import { Col, Container, Row, Form, Button  } from "sveltestrap"
    import GtFormGroupInput from "./../globalUtils/forms/gtFormGroupInput.svelte"
    let config = [{text: "Id:",name: "key", type:"text", id:"key", placeholder:"key"},
        {text: "Name:",name: "name", type:"text", id:"name", placeholder:"name"},
        {text: "Valid Date:", name: "validDate", type:"date", id:"validDate", placeholder:"Valid Date"}
    ]

</script>

<Container>
    <Row>
        <Col class="col-3">
            <Form>
                <GtFormGroupInput {config}/>
                <Button type="submit">Submit</Button>
                <Button type="reset">Clear</Button>
            </Form>
        </Col>
        <Col>
            Products list goes here
        </Col>
    </Row>
</Container>