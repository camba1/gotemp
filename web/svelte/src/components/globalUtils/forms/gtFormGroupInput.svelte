<script>
    import { FormGroup, Input, Label  } from "sveltestrap"
    export let searchParams
</script>

{#each searchParams as searchParamsItem}
    <FormGroup>
        <Label for={searchParamsItem.name}>{searchParamsItem.text}</Label>
        <Input
                type={searchParamsItem.type}
                name={searchParamsItem.name}
                id={searchParamsItem.id}
                placeholder={searchParamsItem.placeholder} />
    </FormGroup>
{/each}