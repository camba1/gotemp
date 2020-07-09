<script>
    export let tblData;
</script>

{#each tblData as dataRow}
    <tr>
        <th scope="row">{dataRow._key}</th>
        <td>{dataRow.name}</td>
        <td>{dataRow.validFrom}</td>
        <td>{dataRow.validThru}</td>
        <td>{dataRow.hierarchylevel}</td>
        <td>{dataRow.otherfields}</td>
    </tr>
{/each}