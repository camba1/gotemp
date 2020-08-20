<script>
    import FormGroup from 'sveltestrap/src/FormGroup.svelte'
    import Input from 'sveltestrap/src/Input.svelte'
    import Label from 'sveltestrap/src/Label.svelte'
    export let searchParams
</script>

{#each searchParams as searchParamsItem}
    <FormGroup>
        <Label for={searchParamsItem.name}>{searchParamsItem.text}</Label>
        <Input
                type={searchParamsItem.type}
                name={searchParamsItem.name}
                id={searchParamsItem.id}
                placeholder={searchParamsItem.placeholder}
                readonly={false}
                bind:value={searchParamsItem.value} />
    </FormGroup>
{/each}
