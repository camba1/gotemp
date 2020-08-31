<script context="module">

    import { httpGet } from '../../globalUtils/api'

    //Helper utils
    import { isObjectEmpty, convertExtraFields, addBoolField } from '../../globalUtils/helperUtils'
    import { userAddresses } from '../../globalUtils/addresses'

    /**
     * Uris for interacting with the server and navigating
     */
    let addresses = userAddresses

    /**
     * String representing the object displayed in the page
     */
    let pageObjectLbl = 'User'

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
                }
            } else {
                alert(`Error getting ${pageObjectLbl}`)
                this.redirect(302, addresses.previousPage);
            }
            let detailData = data
console.log(data)
            return  {detailData, slug, extraFields, addresses} ;

        }

    }

    /**
     * By default grpc will not send out empty fields, so our object may be missing some fields.
     * Additionally, GO will not unMarshal a string value to a boolean field.
     * So we need to initialize the boolean fields with a boolean value so that it is not
     * automatically marshalled as a string
     * @param objToCheck - Object to which we may need to add the boolean field
     * @param flagName - Name of the boolean field to add
     */
    // function addBoolField(objToCheck, flagName) {
    //     if (objToCheck && !objToCheck.hasOwnProperty(flagName)) {
    //         objToCheck[flagName] = false
    //     }
    // }

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

</script>

<Detail {detailData} {slug} {extraFields} {addresses}/>