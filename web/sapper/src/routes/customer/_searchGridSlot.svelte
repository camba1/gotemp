<script>
    /**
     * Component to be displayed within the search grid
     */

    /**
     * Data to be displayed
     * @type {object}
     */
    export let tblData;
</script>

{#each tblData as dataRow}
    <tr>
        <th scope="row"><a href="customer/{dataRow._key}"> {dataRow._key} </a></th>
        <td>{dataRow.name}</td>
        <td>{new Date(dataRow.validityDates.validFrom).toLocaleDateString("en-CA",{year:"numeric",month:"2-digit", day:"2-digit"})}</td>
        <td>{new Date(dataRow.validityDates.validThru).toLocaleDateString("en-CA",{year:"numeric",month:"2-digit", day:"2-digit"})}</td>
    </tr>
{/each}