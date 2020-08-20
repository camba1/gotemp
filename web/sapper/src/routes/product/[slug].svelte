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
    import { httpPut } from '../../globalUtils/api'
    import { httpDelete } from '../../globalUtils/api'
    import GtDetailCard from './../../components/detailScreen/gtDetailCard.svelte'
    import GtDetailCardFormGrp from './../../components/detailScreen/gtDetailCardFormGrp.svelte'
    import Row from 'sveltestrap/src/Row.svelte'
    import Col from 'sveltestrap/src/Col.svelte'
    import Button from 'sveltestrap/src/Button.svelte'
    import Input from 'sveltestrap/src/Input.svelte'

    export let product;
    export let slug;
    export let extraFields;

    let tmpDateFrom = new Date(product.validityDates.validFrom).toLocaleDateString();
    let tmpDateThru = new Date(product.validityDates.validThru).toLocaleDateString();

    async function  handleSave() {
         alert('Saving...');
         console.log(product)
        const {ok, data} = await httpPut("product/productSrv/UpdateProduct", product);
         console.log(ok)
        if (ok) {
            if (isObjectEmpty(data)) {
                alert('Product not saved')
            }
        } else {
            alert('Product not saved')
        }

    }

    async function handleDelete() {
        alert('Deleting...')
        console.log(slug)
        console.log(product)
        const paramString = new URLSearchParams({ _key: `"${slug}"` })
        const {ok, data} = await httpDelete(`product/productSrv/DeleteProduct?${paramString.toString()}`);
        console.log(data)
        if (ok) {
            alert('Product deleted')
        } else {
            alert('Product not deleted')
        }

    }

</script>

{#if product}

    <Row>
        <Col class="col-8">
            <h4>Product: {product.name}</h4>
        </Col>
        <Col class="text-right">
            <Button size="sm" on:click={handleSave}><span><i class="fas fa-save"></i> Save</span></Button>
            <Button size="sm" on:click={handleDelete}><span><i class="fas fa-trash-alt"></i> Delete</span></Button>
            <Button size="sm" onclick="history.back()"><span><i class="fas fa-arrow-alt-circle-left"></i> Back</span></Button>

        </Col>
    </Row>

    <Row>
        <GtDetailCard cardHeader="Information">
            <GtDetailCardFormGrp lblFor="id" lblText="Id:">
                <Input id="id" class="form-control form-control-sm" name="id" type="text" readonly bind:value={product._key}/>
            </GtDetailCardFormGrp>
            <GtDetailCardFormGrp lblFor="name" lblText="Name:">
                <Input id="name"class="form-control form-control-sm"  name="name" type="text" readonly={false} bind:value={product.name}/>
            </GtDetailCardFormGrp>
            <GtDetailCardFormGrp lblFor="hierLevel" lblText="Level:">
                <Input id="hierLevel"class="form-control form-control-sm"  name="hierLevel" type="text" readonly={false} bind:value={product.hierarchyLevel}/>
            </GtDetailCardFormGrp>
        </GtDetailCard>

        <GtDetailCard cardHeader="Validity">
            <GtDetailCardFormGrp lblFor="validFrom" lblText="Valid From:">
                <Input id="validFrom" class="form-control form-control-sm" name="validFrom" type="date" readonly={false}  bind:value={tmpDateFrom}/>
            </GtDetailCardFormGrp>
            <GtDetailCardFormGrp lblFor="validThru" lblText="Valid Thru:">
                <Input id="validThru"class="form-control form-control-sm"  name="validThru" type="date" readonly={false}  bind:value={tmpDateThru}/>
            </GtDetailCardFormGrp>
        </GtDetailCard>
    </Row>
    {#if extraFields}
        <Row>
            <GtDetailCard cardHeader="Other Fields">
                {#each extraFields as field }
                    <GtDetailCardFormGrp lblFor="test2" lblText="{field.name}:">
                        <Input id="test2" class="form-control form-control-sm" name="test2" plaintext readonly={false}  bind:value={field.value}/>
                    </GtDetailCardFormGrp>
                {/each}
            </GtDetailCard>
        </Row>
    {/if}
{:else}
    <h3>No data found for product id: {slug}</h3>
    <Button size="sm" onclick="history.back()"><span><i class="fas fa-arrow-alt-circle-left"></i> Back</span></Button>
{/if}