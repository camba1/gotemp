<script context="module">
    import { authAddresses } from '../../globalUtils/addresses'
    /**
     * Redirect to root page if user is already logged in
     * @param page - is a { host, path, params, query } object
     * @param session - contains user information if logged in
     */
    export async function preload(page, session) {
        if (session.user !== undefined) {
            this.redirect(302, authAddresses.previousPage);
        }
    }
</script>

<script>

    // GUI components imports
    import GtDetailCard from './../../components/detailScreen/gtDetailCard.svelte'
    import GtDetailCardFormGrp from './../../components/detailScreen/gtDetailCardFormGrp.svelte'
    import GTDetailHeader from './../../components/detailScreen/gtDetailHeader.svelte'
    import GTErrorList from './../../components/gtErrorList.svelte'

    import Row from 'sveltestrap/src/Row.svelte'
    import Container from 'sveltestrap/src/Container.svelte'
    import Input from 'sveltestrap/src/Input.svelte'

    // import { authAddresses } from '../../globalUtils/addresses'

    // Allow navigation and Import session to determine if user is logged in
    import { goto, stores } from '@sapper/app'
    const { session } = stores()

    // Interacting with server
    import { httpPost } from '../../globalUtils/api'

    //helper utils
    import { isObjectEmpty } from '../../globalUtils/helperUtils'

    // Default dates
    const curDate = new Date()
    let nextYearDate = new Date()
    const monthsValid = 12
    nextYearDate.setMonth( nextYearDate.getMonth() + 1)

    /**
     * Main object to be displayed in the page
     * @type {object}
     */
    let detailData = {id: 0, firstname: "", lastname: "",
        active: false, pwd:"", name: "",
        email: "",company:"",
        validFrom: curDate, validThru: nextYearDate
    }

    /**
     *  Uris for interacting with the server and navigating
     *  @type {object}
     */
    let addresses = authAddresses

    /**
     * String representing the object displayed in the page
     */
    let pageObjectLbl = 'Registration'

    /**
     * Indicates if an action is in progress and disables the buttons
     * @type {boolean}
     */
    let inProgress = false

    /**
     * list of errors to be displayed to the user. Can include line feeds for multiple lines
     */
    let errorList = null

    /**
     * Handles creation and updates based on whether the slug is present
     * @returns {Promise<void>}
     */
    async function  handleSave() {

        inProgress = true

        if (detailData.name !== "") {
            alert("invalid registration")
            goto(addresses.previousPage)
        } else {

            const {ok, data} = await (httpPost(addresses.register, detailData))
            if (ok) {
                if (isObjectEmpty(data)) {
                    alert(`No data found for ${pageObjectLbl}`)
                } else {
                    goto(addresses.loginPage)
                }
            } else {
                errorList = data
            }

        }
        inProgress = false
    }

    /**
     * Navigates back to main screen
     * @returns {Promise<void>}
     */
    async function backToSearch() {
        await goto(addresses.previousPage)
    }


</script>

<style>
    .dispnon{display: none}
</style>

<Container>

        <Row>
            <GTDetailHeader label="{pageObjectLbl}" inProgress={inProgress} name="" slug=null
                            on:handleSave={handleSave}
                            on:backToSearch={backToSearch} />
        </Row>

        <GTErrorList errorList={errorList} />

        <Row>
            <GtDetailCard cardHeader="New User">
                <GtDetailCardFormGrp lblFor="firstname" lblText="First Name:">
                    <Input id="firstname" class="form-control form-control-sm"  name="firstname" type="text" readonly={false} bind:value={detailData.firstname}/>
                </GtDetailCardFormGrp>
                <GtDetailCardFormGrp lblFor="lastname" lblText="last Name:">
                    <Input id="lastname" class="form-control form-control-sm"  name="lastname" type="text" readonly={false} bind:value={detailData.lastname}/>
                </GtDetailCardFormGrp>
                <GtDetailCardFormGrp lblFor="email" lblText="Email:">
                    <Input id="email" class="form-control form-control-sm"  name="email" type="email" readonly={false} bind:value={detailData.email}/>
                </GtDetailCardFormGrp>
                <GtDetailCardFormGrp lblFor="pwd" lblText="Password:">
                    <Input id="pwd" class="form-control form-control-sm"  name="pwd" type="password" readonly={false} bind:value={detailData.pwd}/>
                </GtDetailCardFormGrp>
            </GtDetailCard>
        </Row>
        <input id="name" class="dispnon" name="name" type="text" bind:value={detailData.name}/>

</Container>