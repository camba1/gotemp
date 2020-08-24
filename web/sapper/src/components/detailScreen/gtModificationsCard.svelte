<script>
    import GtDetailCard from './../../components/detailScreen/gtDetailCard.svelte'
    import GtDetailCardFormGrp from './../../components/detailScreen/gtDetailCardFormGrp.svelte'
    import Input from 'sveltestrap/src/Input.svelte'

    export let createDateTime
    export let updateDateTime
    export let modifiedBy

    let tmpCreateDateTime = new Date(createDateTime).toLocaleString();
    let tmpUpdateDateTime = new Date(updateDateTime).toLocaleString();

</script>

<GtDetailCard cardHeader="Modifications">
    <GtDetailCardFormGrp lblFor="createDate" lblText="Creation Date:">
        <Input id="createDate" class="form-control form-control-sm" name="createDate" plaintext readonly={false}  bind:value={tmpCreateDateTime}/>
    </GtDetailCardFormGrp>
    <GtDetailCardFormGrp lblFor="updateDate" lblText="Update Date:">
        <Input id="updateDate" class="form-control form-control-sm" name="updateDate" plaintext readonly={false}  bind:value={tmpUpdateDateTime}/>
    </GtDetailCardFormGrp>
    <GtDetailCardFormGrp lblFor="modifiedBy" lblText="Modified By:">
        <Input id="modifiedBy" class="form-control form-control-sm" name="modifiedBy" plaintext readonly={false}  bind:value={modifiedBy}/>
    </GtDetailCardFormGrp>
</GtDetailCard>