<script context="module">

    import {httpGet, httpPost} from '../../globalUtils/api'

    //Helper utils
    import { isObjectEmpty, convertExtraFields, addBoolField, addFieldtoObj } from '../../globalUtils/helperUtils'
    import {customerAddresses, promotionAddresses} from '../../globalUtils/addresses'

    /**
     * Uris for interacting with the server and navigating
     */
    let addresses = promotionAddresses

    /**
     * String representing the object displayed in the page
     */
    let pageObjectLbl = 'Promotion'

    /**
     * Redirect to login page if user is not logged in. Load and return page data.
     * Extract the additional fields (that the app is not aware of) into a name, value KV
     * as they will be displayed separately
     * @param page - is a { host, path, params, query } object
     * @param session - contains user information if logged in
     * @returns {Promise<{product: json, slug: *, extraFields: []}>}
     */
    export async function preload(page, session) {
        // console.log(session)
        const { slug } = page.params;

        if (!session.user) {
            this.redirect(302, `/login`);
        } else {

            let extraFields = []
            const paramString = new URLSearchParams({ id: `"${slug}"` })
            const {ok, data} = await httpGet(`${addresses.get}?${paramString.toString()}`, this.fetch, session.token);
            if (ok) {
                // console.log(data)
                if (isObjectEmpty(data)) {

                    alert(`${pageObjectLbl} not found`)
                    this.redirect(302, addresses.new);

                } else {

                    extraFields = convertExtraFields(data.extraFields)
                    addBoolField(data, "active")
                    addFieldtoObj(data, "approvalStatus", 0)
                    addFieldtoObj(data, "prevApprovalStatus", 0)

                }
            } else {

                alert(`Error getting ${pageObjectLbl}`)
                this.redirect(302, addresses.previousPage);

            }
            let detailData = data
            let customersData = await loadCustomers(session)
            // console.log(customersData)
// console.log(data)
            return  {detailData, slug, extraFields, addresses, customersData} ;

        }

    }

    /**
     * Load customer lookup data so that user can pick a customer from a dropdown
     * @param session - session information including session token
     * @returns {Promise<*>} - customer list promise
     */
    async function loadCustomers(session){
            let params = {}
            const {ok, data} = await httpPost(customerAddresses.getAll, params, session.token);
            if (ok) {
                // alert(data)
                if (isObjectEmpty(data)) {
                    alert(`Data not found`)
                } else {
                    // console.log(data)
                    return data.customer
                }
            } else {
                alert(` Error getting data`)
            }
    }

</script>

<script>
    // GUI components imports
    import Detail from './_detail.svelte';

    /**
     * Main object to be displayed in the page
     * @type {object}
     */
    export let detailData;

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
     * customerData holds data for the customer drop down
     * @type {object}
     */
    export let customersData;

</script>
<Detail {detailData} {slug} {extraFields} {addresses} {customersData}/>