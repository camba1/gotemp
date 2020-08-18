<script>
    import FormGroup from 'sveltestrap/src/FormGroup.svelte'
    import Label from 'sveltestrap/src/Label.svelte'
    import Col from 'sveltestrap/src/Col.svelte'

    export let lblFor
    export let lblText

</script>

<FormGroup row>
    <Label for={lblFor} class="col-sm-3 col-form-label col-form-label-sm">{lblText}</Label>
    <Col class="col-sm-9">
        <slot></slot>
    </Col>
</FormGroup>