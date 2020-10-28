<script>
    import RedIcon from '../components/redIcon.svelte';
    import Col from 'sveltestrap/src/Col.svelte'

    export let data
</script>

<style>
    p {
        margin: 0 auto;
    }
</style>

<Col class="col-sm-6 col-lg-4 mb-3">
    <RedIcon  iconName="{data.iconName}" />
    <h5>{data.header}</h5>
    {#each data.techList as tech}
        <p class="text-muted">{tech}</p>
    {/each}
</Col>