
<script context="module">
    /**
     * Generic section that handles loading of data
     * and/or redirect needed prior to rendering the page
     */

    /**
     * Redirect to login page if user is not logged in
     * @param page - is a { host, path, params, query } object
     * @param session - contains user information if logged in
     */
    export function preload(page, session) {
        // console.log(session)
        // console.log(page)
        if (!session.user) {
            this.redirect(302, `/login`);
        }
    }
</script>

<script>
    /**
     * Search page
     */

    // GUI components imports
    import SearchScreen from './../../components/searchScreen/searchScreen.svelte'
    import SearchGridSlot from "./_searchGridSlot.svelte"

    // http Post
    import { httpPost } from '../../globalUtils/api'

    //Helper utils
    import { isObjectEmpty, isValidStringDate } from '../../globalUtils/helperUtils'
    import { productAddresses } from '../../globalUtils/addresses'

    // Allow navigation and Import session to determine if user is logged in
    import { goto, stores } from '@sapper/app'
    const { session } = stores()

    // let tblData=[{"_key": "switch", "name": "Play Switch Console","validityDates":{ "validFrom": "2020-02-02", "validThru": "2021-02-02"}, "hierarchyLevel": "sku", "extraFields": {"externalId": "12345"}},
    //     {"_key": "tele", "name": "Watch me TV", "validityDates":{ "validFrom": "2020-02-02", "validThru": "2021-02-02"}, "hierarchyLevel": "sku", "extraFields": {"externalId": "12345"}},
    //     {"_key": "fridge", "name": "Cool Stuff Fridge", "validityDates":{ "validFrom": "2020-02-02", "validThru": "2021-02-02"}, "hierarchyLevel": "sku", "extraFields": {"externalId": "12345"}}
    // ]

    /**
     * Array of objects to be displayed on the search grid
     * @type {*[]}
     */
    let tblData=[]

    /**
     * Test for search grid column headers
     * @type {string[]}
     */
    let tblHeaders=["ID","name","Valid From","Valid Thru","Hierarchy Level"]

    /**
     * Parameters to be used as search parameters for searching
     * @type {({name: string, text: string, id: string, placeholder: string, type: string, value: string})[]}
     */
    let searchParams = [{text: "Id:",name: "key", type:"text", id:"key", placeholder:"key", value: ""},
        {text: "Name:",name: "name", type:"text", id:"name", placeholder:"name", value:""},
        {text: "Valid Date:", name: "validDate", type:"date", id:"validDate", placeholder:"Valid Date", value:""}
    ]

    /**
     * Title to be displayed on the top of the page.
     * @type {string}
     */
    let pageTitle = "Products"

    /**
     * Uri for the micro-service that will return data for the search grid
     * @type {string}
     */
    //let getDataAddress = "product/productSrv/GetProducts"

    /**
     * Uri to naviaate to when the new and previous buttons are clicked
     * @type {{new: string, previous: string}}
     */
    //let otherPagesAddress = {new: '/product/new', previous: '/' }

    /**
     * Request search results from the server. Data is loaded into the tblData variable.
     * The actual call comes from a child component
     */
    async function handleSearch(event) {

        let params = {_key: searchParams[0].value, name: searchParams[1].value}
        if (isValidStringDate(searchParams[2].value)) {
            params.validDate = new Date(searchParams[2].value).toISOString()
        }
        const {ok, data} = await httpPost(productAddresses.getAll, params, $session.token);
        if (ok) {
            // console.log(data)
            if (isObjectEmpty(data)) {
                alert(`${pageTitle} not found`)
            } else {
                tblData = data.product
            }
        } else {
            alert(`${pageTitle} not found`)
        }
    }

    /**
     * Handle page navigation using sapper's goto method
     * @param newPage - Uri of page we want to visit
     * @returns {Promise<void>}
     */
    async function navigateTo(newPage) {
        switch (newPage.detail.newPage) {
            case "new":
                await goto(productAddresses.new)
                break;
            case "previous":
                await goto(productAddresses.searchPreviousPage);
                break;
            default:
                alert(`Unknown page when trying to navigate: ${newPage.detail.newPage}`);

        }
    }

</script>

<SearchScreen {tblHeaders} {searchParams} {pageTitle} on:search={handleSearch} on:navigate={navigateTo}>
    <SearchGridSlot {tblData}/>
</SearchScreen>
