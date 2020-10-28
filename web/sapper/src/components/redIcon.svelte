<script>
    export let iconName
    let color = "rgb(204,0,0)"
</script>

<style>
    .lnr {
        display: inline-block;
        width: 1em;
        height: 1em;
        vertical-align: -0.05em;
        stroke-width: 1;
    }

    .services-icon {
        margin-bottom: 20px;
        font-size: 30px;
    }
</style>

<span style="color: {color};">
    <i class="fas {iconName} lnr services-icon"></i>
</span>