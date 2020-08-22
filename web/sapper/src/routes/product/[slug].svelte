<script context="module">
    // import { httpPost } from '../../globalUtils/api'
    import { httpGet } from '../../globalUtils/api'

    export async function preload(page, session) {
        const { slug } = page.params;

        let extraFields = []
        // let params = {_key: slug}
        // const {ok, data} = await httpPost("product/productSrv/GetProductById", params);
        const paramString = new URLSearchParams({ _key: `"${slug}"` })
        const {ok, data} = await httpGet(`product/productSrv/GetProductById?${paramString.toString()}`, this.fetch);
        if (ok) {
            console.log(data)
            if (isObjectEmpty(data)) {
                alert('Product not found')
                this.redirect(302, `/product/new`);
            } else {
                extraFields = convertExtraFields(data.extraFields)
            }
        } else {
            // alert('Product not found')
            this.redirect(302, `/product/new`);
        }
        let product = data

        return  {product, slug, extraFields} ;
    }

    function convertExtraFields(obj) {
        let vals = []
        let key;
        if (!isObjectEmpty(obj)) {
            for (let key1 in obj) {
                key = key1
                vals.push({'name': key, 'value': obj[key]})
            }
            return vals
        }
    }

    function isObjectEmpty(obj) {
        for(var i in obj) return false;
        return true;
    }
</script>

<script>

    import Detail from './_detail.svelte';

    export let product;
    export let slug;
    export let extraFields;


</script>

<!--<h5>New Product</h5>-->

<Detail {product} {slug} {extraFields} />