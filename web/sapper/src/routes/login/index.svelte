<script>
    // import Form from 'sveltestrap/src/Form.svelte'
    import Card from 'sveltestrap/src/Card.svelte'
    import CardBody from 'sveltestrap/src/CardBody.svelte'
    import CardHeader from 'sveltestrap/src/CardHeader.svelte'
    import CardTitle from 'sveltestrap/src/CardTitle.svelte'
    import Button from 'sveltestrap/src/Button.svelte'
    import FormGroup from 'sveltestrap/src/FormGroup.svelte'
    import Input from 'sveltestrap/src/Input.svelte'
    import Label from "sveltestrap/src/Label.svelte";
    import { httpPost } from '../../globalUtils/api'
    import { createEventDispatcher } from 'svelte';
    import { goto, stores } from '@sapper/app';

    const dispatch = createEventDispatcher();
    const { session } = stores();

    let email = "duck@mymail.com";
    let pwd = "1234";
    let token = '';

    // function sayHello() {
    //     dispatch('message', {
    //         text: 'Hello!'
    //     });
    // }

    async function login(){
        let user = {email: email, pwd: pwd};
        const {ok, data} = await httpPost("user/userSrv/auth", user);
        if (ok) {
            token = data.token
            $session.user = data.token
            $session.token = data.token
            goto('/');
        } else {
            alert('user not found')
        }
        // console.log(ok);
        // console.log(data);
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
        </form>
    </CardBody>
</Card>
<!--{email}, {pwd} <br> {token}-->