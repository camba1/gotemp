<script>

    /**
     * Component displays non-editable additional fields that are in a NoSQL DB document but that
     * are not part of the application proper
     */

    // GUI components imports
    import GtDetailCard from './../../components/detailScreen/gtDetailCard.svelte'
    import GtDetailCardFormGrp from './../../components/detailScreen/gtDetailCardFormGrp.svelte'
    import Input from 'sveltestrap/src/Input.svelte'

    /**
     * Object that contains the fields to be displayed as key, value pairs
     * @type {object}
     */
    export let extraFields
</script>

<GtDetailCard cardHeader="Other Fields">
    {#each extraFields as field }
        <GtDetailCardFormGrp lblFor="{field.name}" lblText="{field.name}:">
            <Input id="{field.name}" class="form-control-sm" name="{field.name}" plaintext readonly  bind:value={field.value}/>
        </GtDetailCardFormGrp>
    {/each}
</GtDetailCard>