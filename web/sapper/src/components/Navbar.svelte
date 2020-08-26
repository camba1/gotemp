
<script>

    /**
     *  Application navigation bar
     */

    // GUI components imports
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

    // Import session to determine if user is logged in
    import { goto, stores } from '@sapper/app'
    const { session } = stores()

    /**
     * Page segment inidicating which page is displayed in the application currently.
     * Used to control whic link is active in the navigation bar
     * @type {string}
     */
    export let segment

    /**
     * Items to be displayed in the drop down of the navigation bar
     * @type {({label: string, href: string})[]}
     */
    export let servicesDropDown = [{label:"Customers", href:"customer"},
                                    {label:"Products", href:"product"},
                                    {label:"divider", href:""},
                                    {label:"Users", href:"user"},
                                    {label:"blog", href:"blog"},]
    /**
     * Items to be displayed at the top level of the navigation bar (not in the drop dpwn)
     * @type {{label: string, href: string}[]}
     */
    export let servicesTopLevel = [{label:"Promotions", href:"promotion"}]

    /**
     * Indicates if the nav bar toggler is open
     * @type {boolean}
     */
    let isOpen = false;

    /**
     * Set the isOpen varaible
     * @param event
     */
    function handleUpdate(event) {
        isOpen = event.detail.isOpen;
    }

    /**
     * Cleanup the session variables when the user logs out
     * @returns {Promise<void>}
     */
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