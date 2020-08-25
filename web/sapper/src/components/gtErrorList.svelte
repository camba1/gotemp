<script>
    import Row from 'sveltestrap/src/Row.svelte'
    import Col from 'sveltestrap/src/Col.svelte'

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
            <pre>{errorList.detail}</pre>
        </Col>
    </Row>
{/if}