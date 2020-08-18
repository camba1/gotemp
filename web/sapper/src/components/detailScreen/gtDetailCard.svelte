<script>
    import Card from 'sveltestrap/src/Card.svelte'
    import CardBody from 'sveltestrap/src/CardBody.svelte'
    import CardHeader from 'sveltestrap/src/CardHeader.svelte'
    import Col from 'sveltestrap/src/Col.svelte'
    import Form from 'sveltestrap/src/Form.svelte'

    export let cardHeader
</script>
<Col class="col-md-6">
    <Card  class="mb-3">
        <CardHeader>
            {cardHeader}
        </CardHeader>
        <CardBody>
            <Form>
               <slot></slot>
            </Form>
        </CardBody>
    </Card>
</Col>