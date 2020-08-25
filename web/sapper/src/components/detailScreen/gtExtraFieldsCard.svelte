<script>

    import GtDetailCard from './../../components/detailScreen/gtDetailCard.svelte'
    import GtDetailCardFormGrp from './../../components/detailScreen/gtDetailCardFormGrp.svelte'
    import Input from 'sveltestrap/src/Input.svelte'

    export let extraFields
</script>

<GtDetailCard cardHeader="Other Fields">
    {#each extraFields as field }
        <GtDetailCardFormGrp lblFor="{field.name}" lblText="{field.name}:">
            <Input id="{field.name}" class="form-control-sm" name="{field.name}" plaintext readonly  bind:value={field.value}/>
        </GtDetailCardFormGrp>
    {/each}
</GtDetailCard>