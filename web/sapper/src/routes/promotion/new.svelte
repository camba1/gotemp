<script context="module">
    import {httpPost} from "../../globalUtils/api";
    import {customerAddresses} from "../../globalUtils/addresses";
    import {isObjectEmpty} from "../../globalUtils/helperUtils";

    /**
     * Redirect to login page if user is not logged in
     * @param page - is a { host, path, params, query } object
     * @param session - contains user information if logged in
     */
    export async function preload(page, session) {
        if (!session.user) {
            this.redirect(302, `/login`);
        } else {
            let customersData = await loadCustomers(session)
            return  {customersData} ;
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
                data.customer.unshift({_key:"", name:""})
                return data.customer
            }
        } else {
            alert(` Error getting data`)
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
    import { promotionAddresses } from '../../globalUtils/addresses'
    let addresses = promotionAddresses

    /**
     * customerData holds data for the customer drop down
     * @type {object}
     */
    export let customersData;

    // Default dates
    const curDate = new Date()
    let nextYearDate = new Date()
    const monthsValid = 12
    nextYearDate.setMonth( nextYearDate.getMonth() + 1)


    let detailData = {id: 0, name: "",
        active: false, customerId: "",
        description: "",
        approvalStatus: 0, prevApprovalStatus: 0,
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

<Detail {detailData} {slug} {extraFields} {addresses} {customersData}/>