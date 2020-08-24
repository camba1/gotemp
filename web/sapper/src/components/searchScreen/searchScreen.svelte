<script>
    import  Col from 'sveltestrap/src/Col.svelte'
    import Container from 'sveltestrap/src/Container.svelte'
    import Row from 'sveltestrap/src/Row.svelte'
    import Button from 'sveltestrap/src/Button.svelte'
    import Table  from "sveltestrap/src/Table.svelte"
    import GtFormGroupInput from "./../forms/gtFormGroupInput.svelte"

    import { createEventDispatcher } from 'svelte';


    export let tblHeaders;
    export let searchParams;
    export let pageTitle;

    let inProgress = false;

    const dispatch = createEventDispatcher();

    function search() {
        inProgress = true;
        dispatch('message', {
            text: 'Hello!'
        });
        inProgress = false;
    }

    async function openNew() {
        dispatch('navigate', {
            newPage: 'new'
        });
    }
    async function backToPrevious() {
       // await goto('/')
        dispatch('navigate', {
            newPage: 'previous'
        });
    }

</script>

<style>
    thead {
        background-color: rgba(0, 0, 0, .03);
    }
    div.detailgrid {
        overflow: scroll;
        max-height: 600px;
    }

</style>

<Container>
    <Row>
        <Col class="col-9">
            <h4>{pageTitle}</h4>
        </Col>
        <Col class="col-3 text-right">
            <Button size="sm" on:click="{openNew}"><span><i class="fas fa-plus"></i> New</span></Button>
            <Button size="sm" on:click="{backToPrevious}"><span><i class="fas fa-arrow-alt-circle-left"></i> Back</span></Button>
        </Col>
    </Row>
    <Row>
        <Col class="col-3">
            <form on:submit|preventDefault={search}>
                <GtFormGroupInput {searchParams}/>
                    <Button size="sm" type="submit" disabled={inProgress} ><i class="fas fa-search"></i> Search</Button>
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