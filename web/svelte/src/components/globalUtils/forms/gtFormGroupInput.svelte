<script>
    import { FormGroup, Input, Label  } from "sveltestrap"
    export let config
</script>

{#each config as configItem}
    <FormGroup>
        <Label for={configItem.name}>{configItem.text}</Label>
        <Input
                type={configItem.type}
                name={configItem.name}
                id={configItem.id}
                placeholder={configItem.placeholder} />
    </FormGroup>
{/each}