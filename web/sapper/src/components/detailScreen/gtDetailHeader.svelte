<script>

    /**
     * The header for a detail page. Includes a title and the save, delete and back buttons
     */

    // GUI components imports
    import Col from 'sveltestrap/src/Col.svelte'
    import Button from 'sveltestrap/src/Button.svelte'

    // Dispatcher to send events to the parent control
    import { createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher();

    /**
     * Label for the type of the objects we are displaying (e.g.Product, customer, etc)
     * @type {string}
     */
    export let label
    /**
     * Name of the object we are displaying
     * @type {string}
     */
    export let name
    /**
     * Controls editability of the buttons in the form
     * @type {boolean}
     */
    export let inProgress
    /**
     * Identifyer for the object we are displaying
     * @type {string}
     */
    export let slug

    /**
     * Captures the save button press event and forwards to parent component
     * @returns {Promise<void>}
     */
    async function handleSave() {
        dispatch('handleSave');
    }

    /**
     * Captures the delete button press event and forwards to parent component
     * @returns {Promise<void>}
     */
    async function handleDelete() {
        dispatch('handleDelete');
    }

    /**
     * Captures the back to search button press event and forwards to parent component
     * @returns {Promise<void>}
     */
    async function backToSearch() {
        dispatch('backToSearch');
    }

</script>

<Col class="col-8">
    <h4>{label}: {name}</h4>
</Col>
<Col class="text-right" >
    <Button size="sm" disabled={inProgress} on:click={handleSave}><span><i class="fas fa-save"></i> Save</span></Button>
    {#if slug}
        <Button size="sm" disabled={inProgress} on:click={handleDelete}><span><i class="fas fa-trash-alt"></i> Delete</span></Button>
    {/if}
    <Button size="sm" disabled={inProgress} on:click="{backToSearch}"><span><i class="fas fa-arrow-alt-circle-left"></i> Back</span></Button>

</Col>