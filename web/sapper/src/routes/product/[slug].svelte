<script context="module">
    import { httpPost } from '../../globalUtils/api'

    export async function preload(page, session) {
        const { slug } = page.params;


        let params = {_key: slug}
        let extraFields = []
        const {ok, data} = await httpPost("product/productSrv/GetProductById", params);
        if (ok) {
            console.log(data)
            if (isObjectEmpty(data)) {
                alert('Products not found')
            } else {
                extraFields = convertExtraFields(data.extraFields)
            }
        } else {
            alert('Products not found')
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
                alert('Products not saved')
            }
        } else {
            alert('Products not saved')
        }

    }

    function handleDelete() {
        alert('Deleting...')
        console.log(slug)
        console.log(product)
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
                <Input id="name"class="form-control form-control-sm"  name="name" type="text" bind:value={product.name}/>
            </GtDetailCardFormGrp>
            <GtDetailCardFormGrp lblFor="hierLevel" lblText="Level:">
                <Input id="hierLevel"class="form-control form-control-sm"  name="hierLevel" type="text" bind:value={product.hierarchyLevel}/>
            </GtDetailCardFormGrp>
        </GtDetailCard>

        <GtDetailCard cardHeader="Validity">
            <GtDetailCardFormGrp lblFor="validFrom" lblText="Valid From:">
                <Input id="validFrom" class="form-control form-control-sm" name="validFrom" type="date"  bind:value={tmpDateFrom}/>
            </GtDetailCardFormGrp>
            <GtDetailCardFormGrp lblFor="validThru" lblText="Valid Thru:">
                <Input id="validThru"class="form-control form-control-sm"  name="validThru" type="date" bind:value={tmpDateThru}/>
            </GtDetailCardFormGrp>
        </GtDetailCard>
    </Row>
    {#if extraFields}
        <Row>
            <GtDetailCard cardHeader="Other Fields">
                {#each extraFields as field }
                    <GtDetailCardFormGrp lblFor="test2" lblText="{field.name}:">
                        <Input id="test2" class="form-control form-control-sm" name="test2" plaintext  bind:value={field.value}/>
                    </GtDetailCardFormGrp>
                {/each}
            </GtDetailCard>
        </Row>
    {/if}
{:else}
    <h3>No data found for product id: {slug}</h3>
    <Button size="sm" onclick="history.back()"><span><i class="fas fa-arrow-alt-circle-left"></i> Back</span></Button>
{/if}