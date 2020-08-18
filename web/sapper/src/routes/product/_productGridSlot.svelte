<script>
    export let tblData;
</script>

{#each tblData as dataRow}
    <tr>
        <th scope="row"><a href="product/{dataRow._key}"> {dataRow._key} </a></th>
        <td>{dataRow.name}</td>
        <td>{new Date(dataRow.validityDates.validFrom).toLocaleDateString()}</td>
        <td>{new Date(dataRow.validityDates.validThru).toLocaleDateString()}</td>
        <td>{dataRow.hierarchyLevel}</td>
<!--        <td>{dataRow.extraFields}</td>-->
    </tr>
{/each}