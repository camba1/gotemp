<script>
    /**
     * Login page component
     */

    // GUI components imports
    import Card from 'sveltestrap/src/Card.svelte'
    import CardBody from 'sveltestrap/src/CardBody.svelte'
    import CardHeader from 'sveltestrap/src/CardHeader.svelte'
    import CardTitle from 'sveltestrap/src/CardTitle.svelte'
    import Button from 'sveltestrap/src/Button.svelte'
    import FormGroup from 'sveltestrap/src/FormGroup.svelte'
    import Input from 'sveltestrap/src/Input.svelte'
    import Label from "sveltestrap/src/Label.svelte";

    // URLs to different pages and services
    import {authAddresses} from '../../globalUtils/addresses'

    // Post
    import { httpPost } from '../../globalUtils/api'

    // Dispatcher to send events to the parent control
    import { createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher();

    // Linking to pages and sessions
    import { goto, stores } from '@sapper/app';
    const { session } = stores();

    /**
     *  Uris for interacting with the server and navigating
     *  @type {object}
     */
    let addresses = authAddresses

    /**
     * Email used to login to the app
     * @type {string}
     */
    let email = "duck@mymail.com";
    /**
     * Password used to login to the app
     * @type {string}
     */
    let pwd = "1234";

    /**
     * Request a token from the authentication service and stores it in the session store
     * @returns {Promise<void>}
     */
    async function login(){
        let user = {email: email, pwd: pwd};
        const {ok, data} = await httpPost(addresses.auth, user);
        if (ok) {
            $session.user = data.token
            $session.token = data.token
            goto(addresses.previousPage);
        } else {
            alert('user not found')
        }
        // console.log(ok);
        // console.log(data);
    }

    async function backToRoot(){
        await goto(addresses.previousPage)
    }

</script>



<Card class="mb-3 w-50 mx-auto">
    <CardHeader>
        <CardTitle>Login</CardTitle>
    </CardHeader>
    <CardBody>
        <form on:submit|preventDefault={login}>
            <FormGroup>
                <Label for="userEmail">Email:</Label>
                <Input
                        type="email"
                        name="userEmail"
                        id="userEmail"
                        placeholder="Type your email"
                        readonly={false}
                        bind:value={email}/>
            </FormGroup>
            <FormGroup>
                <Label for="userPassword">Password:</Label>
                <Input
                        type="password"
                        name="userPassword"
                        id="userPassword"
                        placeholder="Type your password"
                        readonly={false}
                        bind:value={pwd}/>
            </FormGroup>
            <Button type="submit" >Submit</Button>
            <Button on:click="{backToRoot}">Cancel</Button>
        </form>
    </CardBody>
</Card>
