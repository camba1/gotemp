<script context="module">

    import { httpGet } from '../../globalUtils/api'

    //Helper utils
    import { isObjectEmpty, convertExtraFields } from '../../globalUtils/helperUtils'
    import { productAddresses } from '../../globalUtils/addresses'

    /**
     * Uris for interacting with the server and navigating
     */
    let addresses = productAddresses


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
            const paramString = new URLSearchParams({ _key: `"${slug}"` })
            const {ok, data} = await httpGet(`${addresses.get}?${paramString.toString()}`, this.fetch, session.token);
            if (ok) {
                // console.log(data)
                if (isObjectEmpty(data)) {
                    alert('Product not found')
                    this.redirect(302, addresses.new);
                } else {
                    extraFields = convertExtraFields(data.extraFields)
                }
            } else {
                alert('Error getting Product')
                this.redirect(302, addresses.previousPage);
            }
            let product = data

            return  {product, slug, extraFields, addresses} ;

        }

    }

</script>

<script>

    import Detail from './_detail.svelte';

    export let product;
    export let slug;
    export let extraFields;
    export let addresses;

</script>

<Detail {product} {slug} {extraFields} {addresses}/>