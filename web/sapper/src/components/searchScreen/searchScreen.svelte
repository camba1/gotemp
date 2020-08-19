<script>
    import  Col from 'sveltestrap/src/Col.svelte'
    import Container from 'sveltestrap/src/Container.svelte'
    import Row from 'sveltestrap/src/Row.svelte'
    // import Form from 'sveltestrap/src/Form.svelte'
    import Button from 'sveltestrap/src/Button.svelte'
    import Table  from "sveltestrap/src/Table.svelte"
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
    div.detailgrid {
        overflow-y: scroll;
        max-height: 600px;
    }

</style>

<Container>
    <Row>
        <Col class="col-10">
            <h5>{pageTitle}</h5>
        </Col>
        <Col class="col-2 text-right">
            <Button size="sm"><span><i class="fas fa-plus"></i> Add</span></Button>
        </Col>
    </Row>
    <Row>
        <Col class="col-3">
            <form on:submit|preventDefault={search}>
                <GtFormGroupInput {searchParams}/>
                <Button size="sm" type="submit"><i class="fas fa-search"></i> Search</Button>
                <Button size="sm" type="reset"><i class="fas fa-eraser"></i> Clear</Button>
            </form>
        </Col>
        <Col>
            <div class="detailgrid">
                <Table hover bordered size="sm" class="detailgrid">
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
            </div>
        </Col>
    </Row>
</Container>