<script>
    import { Card, CardBody, CardHeader, CardTitle, Button, Form, FormGroup, Input, Label } from "sveltestrap";
    import { createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher();

    let userName = "";
    let pwd = "";

    function sayHello() {
        dispatch('message', {
            text: 'Hello!'
        });
    }

</script>

<Card class="mb-3 w-50 mx-auto">
    <CardHeader>
        <CardTitle>Login</CardTitle>
    </CardHeader>
    <CardBody>
        <Form>
            <FormGroup>
                <Label for="userEmail">Email:</Label>
                <Input
                        type="email"
                        name="userEmail"
                        id="userEmail"
                        placeholder="Type your email" />
            </FormGroup>
            <FormGroup>
                <Label for="userPassword">Password:</Label>
                <Input
                        type="password"
                        name="userPassword"
                        id="userPassword"
                        placeholder="Type your password" />
            </FormGroup>
            <!--<Button color="primary" type="submit" >Submit</Button>-->
            <Button type="submit" on:click={sayHello} >Submit</Button>
        </Form>
    </CardBody>
</Card>