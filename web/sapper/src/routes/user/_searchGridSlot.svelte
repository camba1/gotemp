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
        <th scope="row"><a href="user/{dataRow.id}"> {dataRow.id} </a></th>
        <td>{dataRow.name}</td>
        <td>{dataRow.company}</td>
        <td>{dataRow.email}</td>
        {#if dataRow.active}
            <td class="text-center"><i class="far fa-check-circle"></i></td>
        {:else}
            <td></td>
        {/if}
        <td>{new Date(dataRow.validFrom).toLocaleDateString("en-CA",{year:"numeric",month:"2-digit", day:"2-digit"})}</td>
        <td>{new Date(dataRow.validThru).toLocaleDateString("en-CA",{year:"numeric",month:"2-digit", day:"2-digit"})}</td>
    </tr>
{/each}