<script>

    import Detail from './_detail.svelte';

    const curDate = new Date()
    let nextYearDate = new Date()
    const monthsValid = 12
    nextYearDate.setMonth( nextYearDate.getMonth() + 1)
    console.log(nextYearDate)
    let product = {_key: "", name: "", validityDates: {validFrom: curDate, validThru: nextYearDate}}
    let slug = ""
    let extraFields = null


</script>

<!--<h5>New Product</h5>-->

<Detail {product} {slug} {extraFields} />