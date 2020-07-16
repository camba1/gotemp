<script>
    import { httpPost } from '../../globalUtils/api'
    import SearchScreen from "./../globalUtils/searchScreen/searchScreen.svelte"
    import ProductGridSlot from "./productGridSlot.svelte"
    let tblData=[{"_key": "switch", "name": "Play Switch Console", "validFrom": "2020-02-02", "validThru": "2021-02-02", "hierarchylevel": "sku", "otherfields": {"externalId": "12345"}},
        {"_key": "tele", "name": "Watch me TV", "validFrom": "2020-02-02", "validThru": "2021-02-02", "hierarchylevel": "sku", "otherfields": {"externalId": "12345"}},
        {"_key": "fridge", "name": "Cool Stuff Fridge", "validFrom": "2020-02-02", "validThru": "2021-02-02", "hierarchylevel": "sku", "otherfields": {"externalId": "12345"}}
    ]
    let tblHeaders=["ID","name","Valid From","Valid Thru","Hierarchy Level","Other Fields"]
    let searchParams = [{text: "Id:",name: "key", type:"text", id:"key", placeholder:"key", value: ""},
        {text: "Name:",name: "name", type:"text", id:"name", placeholder:"name", value:""},
        {text: "Valid Date:", name: "validDate", type:"date", id:"validDate", placeholder:"Valid Date", value:""}
    ]

    let pageTitle = "Products"

    async function handleMessage(event) {
        alert(event.detail.text);
        alert(searchParams[0].value);

        let params = {_key: searchParams[0].value, name: searchParams[1].value}
        if (isValidStringDate(searchParams[2].value)) {
            params.validDate = new Date(searchParams[2].value).toISOString()
        }
        const {ok, data} = await httpPost("product/productSrv/GetProducts", params);
        if (ok) {
            console.log(data)
            if (isObjectEmpty(data)) {
                alert('Products not found')
            } else {
                tblData = data.product
            }
        } else {
            alert('Products not found')
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

<SearchScreen {tblHeaders} {searchParams} {pageTitle} on:message={handleMessage}>
    <ProductGridSlot {tblData}/>
</SearchScreen>
