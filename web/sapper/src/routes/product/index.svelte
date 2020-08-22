<script>
    import { httpPost } from '../../globalUtils/api'
    import SearchScreen from './../../components/searchScreen/searchScreen.svelte'
    import ProductGridSlot from "./_productGridSlot.svelte"

    import { goto } from '@sapper/app'

    // let tblData=[{"_key": "switch", "name": "Play Switch Console","validityDates":{ "validFrom": "2020-02-02", "validThru": "2021-02-02"}, "hierarchyLevel": "sku", "extraFields": {"externalId": "12345"}},
    //     {"_key": "tele", "name": "Watch me TV", "validityDates":{ "validFrom": "2020-02-02", "validThru": "2021-02-02"}, "hierarchyLevel": "sku", "extraFields": {"externalId": "12345"}},
    //     {"_key": "fridge", "name": "Cool Stuff Fridge", "validityDates":{ "validFrom": "2020-02-02", "validThru": "2021-02-02"}, "hierarchyLevel": "sku", "extraFields": {"externalId": "12345"}}
    // ]
    let tblData=[]
    let tblHeaders=["ID","name","Valid From","Valid Thru","Hierarchy Level"] //,"Other Fields"]
    let searchParams = [{text: "Id:",name: "key", type:"text", id:"key", placeholder:"key", value: ""},
        {text: "Name:",name: "name", type:"text", id:"name", placeholder:"name", value:""},
        {text: "Valid Date:", name: "validDate", type:"date", id:"validDate", placeholder:"Valid Date", value:""}
    ]

    let pageTitle = "Products"
    let getDataAddress = "product/productSrv/GetProducts"
    let otherPagesAddress = {new: '/product/new', previous: '/' }

    async function handleMessage(event) {

        let params = {_key: searchParams[0].value, name: searchParams[1].value}
        if (isValidStringDate(searchParams[2].value)) {
            params.validDate = new Date(searchParams[2].value).toISOString()
        }
        const {ok, data} = await httpPost(getDataAddress, params);
        if (ok) {
            console.log(data)
            if (isObjectEmpty(data)) {
                alert(`${pageTitle} not found`)
            } else {
                tblData = data.product
            }
        } else {
            alert(`${pageTitle} not found`)
        }
    }

    async function navigateTo(newPage) {
        switch (newPage.detail.newPage) {
            case "new":
              //  await goto('/product/new');
                await goto(otherPagesAddress.new)
                break;
            case "previous":
                // await goto('/');
                await goto(otherPagesAddress.previous);
                break;
            default:
                alert(`Unknown page when trying to navigate: ${newPage.detail.newPage}`);

        }
    }

    function isObjectEmpty(obj) {
        for(var i in obj) return false;
        return true;
    }

    function isValidDate(d) {
        return d instanceof Date && !isNaN(d);
    }
    function isValidStringDate(stringDate) {
        if (stringDate === "") {
            return false
        }
        let d = new Date(stringDate)
        return isValidDate(d)
    }

</script>

<SearchScreen {tblHeaders} {searchParams} {pageTitle} on:message={handleMessage} on:navigate={navigateTo}>
    <ProductGridSlot {tblData}/>
</SearchScreen>
