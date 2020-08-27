<script>

    // GUI components imports
    import GtDetailCard from './../../components/detailScreen/gtDetailCard.svelte'
    import GtDetailCardFormGrp from './../../components/detailScreen/gtDetailCardFormGrp.svelte'
    import GTModificationsCard from './../../components/detailScreen/gtModificationsCard.svelte'
    import GTExtraFieldsCard from './../../components/detailScreen/gtExtraFieldsCard.svelte'
    import GTDetailHeader from './../../components/detailScreen/gtDetailHeader.svelte'
    import GTErrorList from './../../components/gtErrorList.svelte'
    import GTSaveWarningModal from './../../components/detailScreen/gtSaveWarningModal.svelte'

    import Row from 'sveltestrap/src/Row.svelte'
    import Container from 'sveltestrap/src/Container.svelte'
    import Button from 'sveltestrap/src/Button.svelte'
    import Input from 'sveltestrap/src/Input.svelte'

    // Allow navigation and Import session to determine if user is logged in
    import { goto, stores } from '@sapper/app'
    const { session } = stores()

    // Interacting with server
    import { httpPut, httpPost, httpDelete  } from '../../globalUtils/api'

    //helper utils
    import { isObjectEmpty, updateValidDate } from '../../globalUtils/helperUtils'

    /**
     * Main object to be displayed in the page
     * @type {object}
     */
    export let product;

    /**
     * Slug indicating which item we are displaying (the id)
     * @type {string}
     */
    export let slug;

    /**
     * Field holding any additional values that we know nothing about in the application
     * This is specially valid for NoSQL backends
     * @type ({name: string, value: string}[])
     */
    export let extraFields;

    /**
     *  Uris for interacting with the server and navigating
     *  @type {object}
     */
    export let addresses;

    /**
     * String representing the object displayed in the page
     */
    let pageObjectLbl = 'Product'

    /**
     * String formatted valid from
     * @type {string}
     */
    let tmpDateFrom = product.validityDates && new Date(product.validityDates.validFrom).toLocaleDateString("en-CA",{year:"numeric",month:"2-digit", day:"2-digit"});

    /**
     * String formatted valid thru
     * @type {string}
     */
    let tmpDateThru = product.validityDates && new Date(product.validityDates.validThru).toLocaleDateString("en-CA",{year:"numeric",month:"2-digit", day:"2-digit"});

    /**
     * String formatted create date
     * @type {string}
     */
    let tmpCreateDateTime = product.modifications && new Date(product.modifications.createDate).toLocaleString("en-CA");

    /**
     * String formatted update date
     * @type {string}
     */
    let tmpUpdateDateTime = product.modifications && new Date(product.modifications.updateDate).toLocaleString("en-CA");

    /**
     * Indicates if an action is in progress and disables the buttons
     * @type {boolean}
     */
    let inProgress = false

    /**
     * list of errors to be displayed to the user. Can include line feeds for multiple lines
     */
    let errorList = null

    /**
     * List of warning messages to display to the user. Can include line feeds formultiple items
     */
    let warningMessage = null

    /**
     * Controls opening and closing of the modal that shouw warnings
     */
    let openModal = false

    /**
     * Handles creation and updates based on wether the slug is present
     * @returns {Promise<void>}
     */
    async function  handleSave() {

        inProgress = true
        openModal = false

        const {ok, data} = await (slug
                                    ? httpPut(addresses.update, product, $session.token)
                                    : httpPost(addresses.create, product, $session.token))
        if (ok) {
            if (isObjectEmpty(data)) {
                alert(`No data found for ${pageObjectLbl}`)
            } else if (!slug) {
                goto(addresses.reload + data.product._key)
            } else {
                errorList = null
                product = data.product
                tmpCreateDateTime = new Date(product.modifications.createDate).toLocaleString();
                tmpUpdateDateTime = new Date(product.modifications.updateDate).toLocaleString();
                if (!isObjectEmpty(data.validationErr)) {
                    warningMessage = data.validationErr.failureDesc
                    openModal = true
                }
            }
        } else {
            errorList = data
        }

        inProgress = false
    }

    /**
     * Handles object deletion
     * @returns {Promise<void>}
     */
    async function handleDelete() {
        inProgress = true

        const paramString = new URLSearchParams({ _key: `"${slug}"` })
        const {ok, data} = await httpDelete(`${addresses.delete}?${paramString.toString()}`,null ,$session.token);
        console.log(data)
        if (ok) {
            await backToSearch()
        } else {
            alert(`${pageObjectLbl} not deleted`)
            errorList = data
        }

        inProgress = false
    }

    /**
     * Navigates back to search screen
     * @returns {Promise<void>}
     */
    async function backToSearch() {
        await goto(addresses.previousPage)
    }

    /**
     * Update the valid from data in the page object
     * @param event - item that called the event
     */
    function updateVF(event) {
        updateValidDate("validFrom", event.target.value, product)
    }

    /**
     * Update the valid thru data in the page object
     * @param event - item that called the event
     */
    function updateVT(event) {
        updateValidDate("validThru", event.target.value, product)
    }

</script>

<Container>

    {#if product}

        <Row>
            <GTDetailHeader label="{pageObjectLbl}" inProgress={inProgress} name={product.name} slug={slug}
                            on:handleSave={handleSave}
                            on:handleDelete={handleDelete}
                            on:backToSearch={backToSearch} />
        </Row>

        <GTErrorList errorList={errorList} />
        <GTSaveWarningModal open={openModal} warningText={warningMessage} />

        <Row>
            <GtDetailCard cardHeader="Information">
                <GtDetailCardFormGrp lblFor="id" lblText="Id:">
                    <Input id="id" class="form-control form-control-sm" name="id" type="text" readonly bind:value={product._key}/>
                </GtDetailCardFormGrp>
                <GtDetailCardFormGrp lblFor="name" lblText="Name:">
                    <Input id="name" class="form-control form-control-sm"  name="name" type="text" readonly={false} bind:value={product.name}/>
                </GtDetailCardFormGrp>
                <GtDetailCardFormGrp lblFor="hierLevel" lblText="Level:">
                    <Input id="hierLevel" class="form-control form-control-sm"  name="hierLevel" type="text" readonly={false} bind:value={product.hierarchyLevel}/>
                </GtDetailCardFormGrp>
            </GtDetailCard>

            <GtDetailCard cardHeader="Validity">
                <GtDetailCardFormGrp lblFor="validFrom" lblText="Valid From:">
                    <Input id="validFrom" class="form-control form-control-sm" name="validFrom" type="date" placeholder="yyyy-mm-dd" readonly={false} on:input={updateVF} bind:value={tmpDateFrom}/>
                </GtDetailCardFormGrp>
                <GtDetailCardFormGrp lblFor="validThru" lblText="Valid Thru:">
                    <Input id="validThru" class="form-control form-control-sm"  name="validThru" type="date" placeholder="yyyy-mm-dd" readonly={false} on:input={updateVT} bind:value={tmpDateThru}/>
                </GtDetailCardFormGrp>
            </GtDetailCard>
        </Row>

        <Row>
            {#if product.modifications}
                <GTModificationsCard createDateTime={tmpCreateDateTime} updateDateTime={tmpUpdateDateTime} modifiedBy={product.modifications.modifiedBy} />
            {/if}
            {#if extraFields}
                <GTExtraFieldsCard extraFields={extraFields} />
            {/if}
        </Row>

    {:else}
        <h3>No data found for {pageObjectLbl}</h3>
        <Button size="sm" on:click="{backToSearch}"><span><i class="fas fa-arrow-alt-circle-left"></i> Back</span></Button>
    {/if}

</Container>