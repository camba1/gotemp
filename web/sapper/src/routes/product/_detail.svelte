<script>
    import { goto } from '@sapper/app'
    import { httpPut } from '../../globalUtils/api'
    import { httpPost } from '../../globalUtils/api'
    import { httpDelete } from '../../globalUtils/api'

    import GtDetailCard from './../../components/detailScreen/gtDetailCard.svelte'
    import GtDetailCardFormGrp from './../../components/detailScreen/gtDetailCardFormGrp.svelte'
    import GTModificationsCard from './../../components/detailScreen/gtModificationsCard.svelte'
    import GTExtraFieldsCard from './../../components/detailScreen/gtExtraFieldsCard.svelte'
    import GTDetailHeader from './../../components/detailScreen/gtDetailHeader.svelte'
    import GTErrorList from './../../components/gtErrorList.svelte'

    import Row from 'sveltestrap/src/Row.svelte'
    import Container from 'sveltestrap/src/Container.svelte'
    import Button from 'sveltestrap/src/Button.svelte'
    import Input from 'sveltestrap/src/Input.svelte'

    export let product;
    export let slug;
    export let extraFields;

    let tmpDateFrom = new Date(product.validityDates.validFrom).toLocaleDateString("en-CA",{year:"numeric",month:"2-digit", day:"2-digit"});
    let tmpDateThru = new Date(product.validityDates.validThru).toLocaleDateString("en-CA",{year:"numeric",month:"2-digit", day:"2-digit"});
    let tmpCreateDateTime = new Date(product.modifications.createDate).toLocaleString("en-CA");
    let tmpUpdateDateTime = new Date(product.modifications.updateDate).toLocaleString("en-CA");

    let addresses = {update: "product/productSrv/UpdateProduct",
        create: "product/productSrv/CreateProduct",
        delete: "product/productSrv/DeleteProduct",
        reload: "product/",
        previousPage: "/product"}

    let inProgress = false

    let errorList = null

    async function  handleSave() {

        inProgress = true

        const {ok, data} = await (slug
                                    ? httpPut(addresses.update, product)
                                    : httpPost(addresses.create, product))
        if (ok) {
            if (isObjectEmpty(data)) {
                alert('No data found for product')
            } else if (!slug) {
                goto(addresses.reload + data.product._key)
            } else {
                errorList = null
                product = data.product
                tmpCreateDateTime = new Date(product.modifications.createDate).toLocaleString();
                tmpUpdateDateTime = new Date(product.modifications.updateDate).toLocaleString();
            }
        } else {
            errorList = data
        }

        inProgress = false
    }

    async function handleDelete() {
        inProgress = true

        const paramString = new URLSearchParams({ _key: `"${slug}"` })
        const {ok, data} = await httpDelete(`${addresses.delete}?${paramString.toString()}`);
        console.log(data)
        if (ok) {
            await backToSearch()
        } else {
            alert('Product not deleted')
        }

        inProgress = false
    }

    async function backToSearch() {

        await goto(addresses.previousPage)
    }

    function isObjectEmpty(obj) {
        for(var i in obj) return false;
        return true;
    }

    function updateVF(event) {
        updateVD("validFrom", event.target.value)
    }

    function updateVT(event) {
        updateVD("validThru", event.target.value)
    }

    function updateVD(dateToUpdate, newDateString){
        let foundVD = false
        if (product) {
            if (product.validityDates) {
                let parts = newDateString.split('-');
                let VD = new Date(parts[0], parts[1] - 1, parts[2]);
                // let VD = new Date(event.target.value)
                if (isValidDate(VD)){
                    product.validityDates[dateToUpdate] = VD
                }
                foundVD = true
            }
        }
        if (!foundVD) {
            alert('Unable to populate validity date')
        }
    }

    function isValidDate(d) {
        return d instanceof Date && !isNaN(d);
    }

</script>

<Container>

    {#if product}

        <Row>
            <GTDetailHeader label="Product" inProgress={inProgress} name={product.name} slug={slug}
                            on:handleSave={handleSave}
                            on:handleDelete={handleDelete}
                            on:backToSearch={backToSearch} />
        </Row>


        <GTErrorList errorList={errorList} />

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
                    <Input id="validFrom" class="form-control form-control-sm" name="validFrom" type="date" placeholder="yyyy-mm-dd" readonly={false} on:input={updateVF} bind:value={tmpDateFrom}/>
                </GtDetailCardFormGrp>
                <GtDetailCardFormGrp lblFor="validThru" lblText="Valid Thru:">
                    <Input id="validThru"class="form-control form-control-sm"  name="validThru" type="date" placeholder="yyyy-mm-dd" readonly={false} on:input={updateVT} bind:value={tmpDateThru}/>
                </GtDetailCardFormGrp>
            </GtDetailCard>
        </Row>

        <Row>
            {#if product.modifications}
                <GTModificationsCard createDateTime={tmpCreateDateTime} updateDateTime={tmpUpdateDateTime} modifiedBy={product.modifications.modifiedBy} />
            {/if}
            {#if extraFields}
                <GTExtraFieldsCard extraFields={extraFields} />
            {/if}
        </Row>

    {:else}
        <h3>No data found for product</h3>
        <Button size="sm" on:click="{backToSearch}"><span><i class="fas fa-arrow-alt-circle-left"></i> Back</span></Button>
    {/if}

</Container>