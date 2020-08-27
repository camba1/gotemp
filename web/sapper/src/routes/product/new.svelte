<script context="module">
    export async function preload(page, session) {
        if (!session.user) {
            this.redirect(302, `/login`);
        }
    }
</script>
<script>

    import Detail from './_detail.svelte';
    import { productAddresses } from '../../globalUtils/addresses'
    let addresses = productAddresses

    const curDate = new Date()
    let nextYearDate = new Date()
    const monthsValid = 12
    nextYearDate.setMonth( nextYearDate.getMonth() + 1)
    let product = {_key: "", name: "", validityDates: {validFrom: curDate, validThru: nextYearDate}}
    let slug = ""
    let extraFields = null


</script>

<Detail {product} {slug} {extraFields} {addresses} />