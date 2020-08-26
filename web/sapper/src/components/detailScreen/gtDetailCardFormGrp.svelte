<script>

    /**
     * Contatenated label + input control. Epects that the actual input will be pased in
     * and placed in the slot sction
     */

    // GUI components imports
    import FormGroup from 'sveltestrap/src/FormGroup.svelte'
    import Label from 'sveltestrap/src/Label.svelte'
    import Col from 'sveltestrap/src/Col.svelte'

    /**
     * Indicates which control the label relates to.
     * @type {string}
     */
    export let lblFor
    /**
     * Text to be displayed in the label
     * @type {string}
     */
    export let lblText

</script>

<FormGroup row>
    <Label for={lblFor} class="col-sm-3 col-form-label col-form-label-sm">{lblText}</Label>
    <Col class="col-sm-9">
        <slot></slot>
    </Col>
</FormGroup>