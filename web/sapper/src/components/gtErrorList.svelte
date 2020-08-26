<script>
    /**
     * Component to display errors in the application
     */

    // GUI components imports
    import Row from 'sveltestrap/src/Row.svelte'
    import Col from 'sveltestrap/src/Col.svelte'

    /**
     * list of errors. Can contain line feeds for multiple errors
     * @type {string}
     */
    export let errorList;

</script>

<style>
    pre {
        display: block;
        font-family: Roboto, -apple-system, BlinkMacSystemFont, Segoe UI, Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue, sans-serif;
        font-weight: bold;
        white-space: pre;
        margin: 1em 0;
        color: #8B0000;
    }
</style>
{#if errorList}
    <Row>
        <Col class="text-center" >
            <span style="color: #8B0000"><i class="fas fa-times-circle"></i><pre>{errorList.detail}</pre></span>
        </Col>
    </Row>
{/if}