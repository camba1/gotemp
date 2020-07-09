<script>
    import SearchScreen from "./../globalUtils/searchScreen/searchScreen.svelte"
    import ProductGridSlot from "./productGridSlot.svelte"
    let tblData=[{"_key": "switch", "name": "Play Switch Console", "validFrom": "2020-02-02", "validThru": "2021-02-02", "hierarchylevel": "sku", "otherfields": {"externalId": "12345"}},
        {"_key": "tele", "name": "Watch me TV", "validFrom": "2020-02-02", "validThru": "2021-02-02", "hierarchylevel": "sku", "otherfields": {"externalId": "12345"}},
        {"_key": "fridge", "name": "Cool Stuff Fridge", "validFrom": "2020-02-02", "validThru": "2021-02-02", "hierarchylevel": "sku", "otherfields": {"externalId": "12345"}}
    ]
    let tblHeaders=["ID","name","Valid From","Valid Thru","Hierarchy Level","Other Fields"]
    let searchParams = [{text: "Id:",name: "key", type:"text", id:"key", placeholder:"key"},
        {text: "Name:",name: "name", type:"text", id:"name", placeholder:"name"},
        {text: "Valid Date:", name: "validDate", type:"date", id:"validDate", placeholder:"Valid Date"}
    ]
    let pageTitle = "Products"
</script>

<SearchScreen {tblHeaders} {searchParams} {pageTitle}>
    <ProductGridSlot {tblData}/>
</SearchScreen>
