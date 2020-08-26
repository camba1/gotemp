<script>

    /**
     * Component displays audit fields of an object (creation date, update date , etc)
     */

    // GUI components imports
    import GtDetailCard from './../../components/detailScreen/gtDetailCard.svelte'
    import GtDetailCardFormGrp from './../../components/detailScreen/gtDetailCardFormGrp.svelte'
    import Input from 'sveltestrap/src/Input.svelte'

    /**
     * Creation date and time in the correct locale to be displayed
     * @type {string}
     */
    export let createDateTime
    /**
     * Last update date and time in the correct locale to be displayed
     * @type {string}
     */
    export let updateDateTime
    /**
     * Last user to modify the document/record
     * @type {string}
     */
    export let modifiedBy


</script>

<GtDetailCard cardHeader="Modifications">
    <GtDetailCardFormGrp lblFor="createDate" lblText="Creation Date:">
        <Input id="createDate" class="form-control-sm" name="createDate" plaintext readonly  bind:value={createDateTime}/>
    </GtDetailCardFormGrp>
    <GtDetailCardFormGrp lblFor="updateDate" lblText="Update Date:">
        <Input id="updateDate" class="form-control-sm" name="updateDate" plaintext readonly  bind:value={updateDateTime}/>
    </GtDetailCardFormGrp>
    <GtDetailCardFormGrp lblFor="modifiedBy" lblText="Modified By:">
        <Input id="modifiedBy" class="form-control-sm" name="modifiedBy" plaintext readonly  bind:value={modifiedBy}/>
    </GtDetailCardFormGrp>
</GtDetailCard>