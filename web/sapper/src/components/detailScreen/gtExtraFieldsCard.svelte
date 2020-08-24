<script>

    import GtDetailCard from './../../components/detailScreen/gtDetailCard.svelte'
    import GtDetailCardFormGrp from './../../components/detailScreen/gtDetailCardFormGrp.svelte'
    import Input from 'sveltestrap/src/Input.svelte'

    export let extraFields
</script>

<GtDetailCard cardHeader="Other Fields">
    {#each extraFields as field }
        <GtDetailCardFormGrp lblFor="test2" lblText="{field.name}:">
            <Input id="test2" class="form-control form-control-sm" name="test2" plaintext readonly={false}  bind:value={field.value}/>
        </GtDetailCardFormGrp>
    {/each}
</GtDetailCard>