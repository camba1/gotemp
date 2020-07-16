<script>
    import { Card, CardBody, CardHeader, CardTitle, Button, Form, FormGroup, Input, Label } from "sveltestrap";
    import { httpPost } from '../../globalUtils/api'
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();


    let email; // = "duck@mymail.com";
    let pwd; // = "1234";
    let token = '';

    function sayHello() {
        dispatch('message', {
            text: 'Hello!'
        });
    }

    async function login(){
        let user = {email: email, pwd: pwd};
        const {ok, data} = await httpPost("user/userSrv/auth", user);
        if (ok) {
            token = data.token
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
                        bind:value={email}/>
            </FormGroup>
            <FormGroup>
                <Label for="userPassword">Password:</Label>
                <Input
                        type="password"
                        name="userPassword"
                        id="userPassword"
                        placeholder="Type your password"
                        bind:value={pwd}/>
            </FormGroup>
            <Button type="submit" >Submit</Button>
        </form>
    </CardBody>
</Card>
<!--{email}, {pwd} <br> {token}-->