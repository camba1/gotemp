<script context="module">
    /**
     * Redirect to login page if user is not logged in
     * @param page - is a { host, path, params, query } object
     * @param session - contains user information if logged in
     */
    export async function preload(page, session) {
        if (!session.user) {
            this.redirect(302, `/login`);
        }
    }
</script>
<script>

    /**
     * New Item page
     */

    // GUI components imports
    import Detail from './_detail.svelte';
    // URLs to different pages and product services
    import { userAddresses } from '../../globalUtils/addresses'
    let addresses = userAddresses

    // Default dates
    const curDate = new Date()
    let nextYearDate = new Date()
    const monthsValid = 12
    nextYearDate.setMonth( nextYearDate.getMonth() + 1)


    let detailData = {id: 0, firstname: "", lastname: "",
        active: false, pwd:"", name: "",
        email: "",company:"",
        validFrom: curDate, validThru: nextYearDate
        // validityDates: {validFrom: curDate, validThru: nextYearDate}
    }
    /**
     * Slug to be passed to child component. It should be initially empty
     * @type {string}
     */
    let slug = ""
    /**
     * Extra fields to be populated by child component. It should be initially null
     * @type {null}
     */
    let extraFields = null


</script>

<Detail {detailData} {slug} {extraFields} {addresses} />