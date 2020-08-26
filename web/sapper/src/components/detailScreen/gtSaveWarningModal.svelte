<script>

    /**
     * Component modal that shows warnings that do not stop obejct from being saved
     */

    // GUI components imports
    import Button from 'sveltestrap/src/Button.svelte'
    import Modal from 'sveltestrap/src/Modal.svelte'
    import ModalBody from 'sveltestrap/src/ModalBody.svelte'
    import ModalFooter from 'sveltestrap/src/ModalFooter.svelte'
    import ModalHeader from 'sveltestrap/src/ModalHeader.svelte'

    /**
     * Indicates whether the modal is open and displayed to the user
     * @type {boolean}
     */
    export let open = false
    /**
     * Text to display in the modal
     * @type {string}
     */
    export let warningText
    /**
     * Controls the visibility of the modal
     * @returns {boolean}
     */
    const toggle = () => (open = !open)

</script>

<style>
    pre {
        display: block;
        font-family: Roboto, -apple-system, BlinkMacSystemFont, Segoe UI, Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue, sans-serif;
        white-space: pre;
        margin: 1em 0;
    }
</style>

<div>
    <Modal isOpen={open} {toggle}>
        <ModalHeader class="bg-light" {toggle}><span style="color: darkgoldenrod"><i class="fas fa-exclamation-triangle"></i> Warning</span></ModalHeader>
        <ModalBody>
            <span>Data was saved with warnings:</span>
            <pre>{warningText}</pre>
        </ModalBody>
        <ModalFooter class="bg-light">
            <Button on:click={toggle}>OK</Button>
        </ModalFooter>
    </Modal>
</div>