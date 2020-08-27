
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
    import { isObjectEmpty } from '../../globalUtils/helperUtils'
    import { customerAddresses } from '../../globalUtils/addresses'

    // Allow navigation and Import session to determine if user is logged in
    import { goto, stores } from '@sapper/app'
    const { session } = stores()

    // let tblData=[{"_key": "ducksrus","name": "Ducks R Us", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "externalId":"12345"},
    //              {"_key": "fridge", "name": "Cool Stuff Fridge", "validityDates":{ "validFrom": "2020-02-02", "validThru": "2021-02-02"}, "hierarchyLevel": "sku", "extraFields": {"externalId": "12345"}}
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
    let tblHeaders=["ID","name","Valid From","Valid Thru"]

    /**
     * Parameters to be used as search parameters for searching
     * @type {({name: string, text: string, id: string, placeholder: string, type: string, value: string})[]}
     */
    let searchParams = [{text: "Id:",name: "key", type:"text", id:"key", placeholder:"key", value: ""},
        {text: "Name:",name: "name", type:"text", id:"name", placeholder:"name", value:""}
    ]

    /**
     * Title to be displayed on the top of the page.
     * @type {string}
     */
    let pageTitle = "Customers"


    /**
     * Request search results from the server. Data is loaded into the tblData variable.
     * The actual call comes from a child component
     */
    async function handleSearch(event) {

        let params = {_key: searchParams[0].value, name: searchParams[1].value}
        const {ok, data} = await httpPost(customerAddresses.getAll, params, $session.token);
        if (ok) {
            // console.log(data)
            if (isObjectEmpty(data)) {
                alert(`${pageTitle} not found`)
            } else {
                tblData = data.customer
            }
        } else {
            alert(` Error getting ${pageTitle}`)
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
                await goto(customerAddresses.new)
                break;
            case "previous":
                await goto(customerAddresses.searchPreviousPage);
                break;
            default:
                alert(`Unknown page when trying to navigate: ${newPage.detail.newPage}`);

        }
    }

</script>

<SearchScreen {tblHeaders} {searchParams} {pageTitle} on:search={handleSearch} on:navigate={navigateTo}>
    <SearchGridSlot {tblData}/>
</SearchScreen>
