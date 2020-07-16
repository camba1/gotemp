<script>
    import { Col, Container, Row, Form, Button, Table  } from "sveltestrap"
    import { createEventDispatcher } from 'svelte';
    import GtFormGroupInput from "./../forms/gtFormGroupInput.svelte"

    export let tblHeaders;
    export let searchParams;
    export let pageTitle;

    const dispatch = createEventDispatcher();

    function search() {
        dispatch('message', {
            text: 'Hello!'
        });
    }

</script>

<style>
    thead {
        background-color: rgba(0, 0, 0, .03);
    }
</style>

<Container>
    <h5>{pageTitle}</h5>
    <Row>
        <Col class="col-3">
            <form on:submit|preventDefault={search}>
                <GtFormGroupInput {searchParams}/>
                <Button type="submit">Search</Button>
                <Button type="reset">Clear</Button>
            </form>
        </Col>
        <Col>
            <Table hover bordered size="sm">
                <thead>
                <tr>
                    {#each tblHeaders as header}
                        <th scope="col">{header}</th>
                    {/each}
                </tr>
                </thead>
                <tbody>
                <slot></slot>
                </tbody>
            </Table>
        </Col>
    </Row>
</Container>