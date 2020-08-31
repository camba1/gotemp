
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
    import {isObjectEmpty, isValidStringDate} from '../../globalUtils/helperUtils'
    import { userAddresses } from '../../globalUtils/addresses'

    // Allow navigation and Import session to determine if user is logged in
    import { goto, stores } from '@sapper/app'
    const { session } = stores()

    /**
     * Array of objects to be displayed on the search grid
     * @type {*[]}
     */
    let tblData=[]

    /**
     * Test for search grid column headers
     * @type {string[]}
     */
    let tblHeaders=["ID","Name","Company", "Email", "Active", "Valid From","Valid Thru"]

    /**
     * Parameters to be used as search parameters for searching
     * @type {({name: string, text: string, id: string, placeholder: string, type: string, value: string})[]}
     */
    let searchParams = [{text: "Id:",name: "id", type:"text", id:"id", placeholder:"Id", value: ""},
        {text: "First Name:",name: "firstName", type:"text", id:"firstName", placeholder:"First Name", value:""},
        {text: "Last Name:",name: "lastName", type:"text", id:"lastName", placeholder:"Last Name", value:""},
        {text: "Company:",name: "company", type:"text", id:"company", placeholder:"Company", value:""},
        {text: "Email:",name: "email", type:"email", id:"email", placeholder:"Email address", value:""},
        {text: "Valid Date:", name: "validDate", type:"date", id:"validDate", placeholder:"Valid Date", value:""}

    ]

    /**
     * Title to be displayed on the top of the page.
     * @type {string}
     */
    let pageTitle = "Users"


    /**
     * Request search results from the server. Data is loaded into the tblData variable.
     * The actual call comes from a child component
     */
    async function handleSearch(event) {


        let params = {
            fisrtname: searchParams[1].value,
            lastname: searchParams[2].value,
            company: searchParams[3].value,
            email: searchParams[4].value,
        }
        //Go microservice will not handle empty id field since id is expected to be int value
        // and empty for int values is 0 in Go
        if (searchParams[0].value !== "") {
            !isNaN(searchParams[0].value) ?
                params.id = searchParams[0].value : searchParams[0].value = ""
        }
        //Validate we have a date before adding it to the request
        if (isValidStringDate(searchParams[5].value)) {
            params.validDate = new Date(searchParams[5].value).toISOString()
        }

        const {ok, data} = await httpPost(userAddresses.getAll, params, $session.token);
        if (ok) {
            // console.log(data)
            if (isObjectEmpty(data)) {
                alert(`${pageTitle} not found`)
            } else {
                tblData = data.user
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
                await goto(userAddresses.new)
                break;
            case "previous":
                await goto(userAddresses.searchPreviousPage);
                break;
            default:
                alert(`Unknown page when trying to navigate: ${newPage.detail.newPage}`);

        }
    }

</script>

<SearchScreen {tblHeaders} {searchParams} {pageTitle} on:search={handleSearch} on:navigate={navigateTo}>
    <SearchGridSlot {tblData}/>
</SearchScreen>
