<script>
    import Input from 'sveltestrap/src/Input.svelte'
    import FormGroup from 'sveltestrap/src/FormGroup.svelte'
    import Label from 'sveltestrap/src/Label.svelte'
    import Col from 'sveltestrap/src/Col.svelte'

    export let name
    export let text
    export let id
    export let type
    export let value
    export let placeholder

</script>

<FormGroup row>
<!--    <Label for="name" class="col-sm-2 col-form-label col-form-label-sm">Name:</Label>-->
    <Label for={name} class="col-sm-2 col-form-label col-form-label-sm">{text}</Label>
    <Col class="col-sm-10">
<!--        <Input id="name"class="form-control form-control-sm"  name="name" type="text" bind:value={data.name}/>-->
        <Input  class="form-control form-control-sm"
                type={type}
                name={name}
                id={id}
                placeholder={placeholder}
                bind:value={value} />
    </Col>
</FormGroup>