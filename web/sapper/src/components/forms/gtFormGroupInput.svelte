<script>

    /**
     * While this was built for dispaying search parameters on search screen,
     * it is a generic horizontal form group to display in a form
     */

    // GUI components imports
    import FormGroup from 'sveltestrap/src/FormGroup.svelte'
    import Input from 'sveltestrap/src/Input.svelte'
    import Label from 'sveltestrap/src/Label.svelte'

    /**
     * The items to be displayed in the form group
     * @type {object}
     * @property {string} name
     * @property {string} text
     * @property {string} type
     * @property {string} id
     * @property {string} placeholder
     * @property {string} value
     */
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
