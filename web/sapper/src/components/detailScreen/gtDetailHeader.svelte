<script>
    import Col from 'sveltestrap/src/Col.svelte'
    import Button from 'sveltestrap/src/Button.svelte'

    import { createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher();

    export let label
    export let name
    export let inProgress
    export let slug

    async function handleSave() {
        dispatch('handleSave');
    }

    async function handleDelete() {
        dispatch('handleDelete');
    }

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