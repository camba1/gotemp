<script>
    /**
     * Component shows the search screen with search parameters in one side
     * and a results grid on the other
     */

    // GUI components imports
    import  Col from 'sveltestrap/src/Col.svelte'
    import Container from 'sveltestrap/src/Container.svelte'
    import Row from 'sveltestrap/src/Row.svelte'
    import Button from 'sveltestrap/src/Button.svelte'
    import Table  from "sveltestrap/src/Table.svelte"
    import GtFormGroupInput from "./../forms/gtFormGroupInput.svelte"

    // Dispatcher to send events to the parent control
    import { createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher();

    /**
     * String array of results grid headers
     * @type {array}
     */
    export let tblHeaders;
    /**
     * The items to be used as search parameters
     * @type {object}
     * @property {string} name
     * @property {string} text
     * @property {string} type
     * @property {string} id
     * @property {string} placeholder
     * @property {string} value
     */
    export let searchParams;

    /**
     * Search screen title
     * @type {string}
     */
    export let pageTitle;

    /**
     * Controls editability of buttons on the screen
     * @type {boolean}
     */
    let inProgress = false;


    /**
     * Captures the search button click and forwards to parent component
     */
    function search() {
        inProgress = true;
        dispatch('message', {
            text: 'Hello!'
        });
        inProgress = false;
    }

    /**
     * Captures the new button click and forwards to parent component
     * @returns {Promise<void>}
     */
    async function openNew() {
        dispatch('navigate', {
            newPage: 'new'
        });
    }

    /**
     * Captures the back button click and forwards to parent component
     * @returns {Promise<void>}
     */
    async function backToPrevious() {
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