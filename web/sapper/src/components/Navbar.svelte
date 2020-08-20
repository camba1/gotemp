
<script>

    import Collapse from 'sveltestrap/src/Collapse.svelte'
    import Navbar from 'sveltestrap/src/Navbar.svelte'
    import NavbarToggler from 'sveltestrap/src/NavbarToggler.svelte'
    import NavbarBrand from 'sveltestrap/src/NavbarBrand.svelte'
    import Nav from 'sveltestrap/src/Nav.svelte'
    import NavItem from 'sveltestrap/src/NavItem.svelte'
    import NavLink from 'sveltestrap/src/NavLink.svelte'
    import UncontrolledDropdown from 'sveltestrap/src/UncontrolledDropdown.svelte'
    import DropdownToggle from 'sveltestrap/src/DropdownToggle.svelte'
    import DropdownMenu  from 'sveltestrap/src/DropdownMenu.svelte'
    import DropdownItem from 'sveltestrap/src/DropdownItem.svelte'
    import Button from 'sveltestrap/src/Button.svelte'

    import { goto, stores } from '@sapper/app'
    const { session } = stores()

    // export let isLoggedOn = false
    export let segment

    export let servicesDropDown = [{label:"Customers", href:"customer"},
                                    {label:"Products", href:"product"},
                                    {label:"divider", href:""},
                                    {label:"Users", href:"user"},
                                    {label:"blog", href:"blog"},]
    export let servicesTopLevel = [{label:"Promotions", href:"promotion"}]

    let isOpen = false;

    function handleUpdate(event) {
        isOpen = event.detail.isOpen;
    }

    async function logout(){
        $session.user = undefined
        $session.token = undefined
    }

</script>

<Navbar color="light" fixed="top" light expand="md">
    <NavbarBrand href="/">goTemp</NavbarBrand>
    <NavbarToggler on:click={() => (isOpen = !isOpen)} />
    <Collapse {isOpen} navbar expand="md" on:update={handleUpdate}>
        <Nav class="ml-auto" navbar>
            {#if $session.token != null}
                {#each servicesTopLevel as service}
                    <NavItem>
                        <NavLink href="{service.href}" active="{segment === service.href ? 'page' : undefined}" >{service.label}</NavLink>
                    </NavItem>
                {/each}
                <UncontrolledDropdown nav inNavbar>
                    <DropdownToggle nav caret>Services</DropdownToggle>
                    <DropdownMenu right>
                        {#each servicesDropDown as service}
                            {#if service.href === "" }
                                <DropdownItem divider />
                            {:else}
                                <DropdownItem><NavLink href="{service.href}" active="{segment === service.href ? 'page' : undefined}">{service.label}</NavLink></DropdownItem>
                            {/if}
                        {/each}
                    </DropdownMenu>
                </UncontrolledDropdown>
                <NavItem>
                    <NavLink on:click={logout}><span><i class="fas fa-sign-out-alt"></i></span></NavLink>
                </NavItem>
            {:else}
                <NavItem>
                    <NavLink href="login" active="{segment === 'login' ? 'page' : undefined}">Login</NavLink>
                </NavItem>
                <NavItem>
                    <NavLink href="register" active="{segment === 'register' ? 'page' : undefined}">Register</NavLink>
                </NavItem>
            {/if}

            <NavItem>
                <NavLink href="https://bitbucket.org/Bolbeck/gotemp"><i class="fab fa-bitbucket"></i></NavLink>
            </NavItem>
        </Nav>
    </Collapse>
</Navbar>