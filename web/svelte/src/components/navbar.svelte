<!--<script>-->
<!--    import Navbar from 'sveltestrap/src/Navbar.svelte';-->
<!--    export let navList = [];-->
<!--    export let header ;-->
<!--</script>-->

<script>

    import {
        Collapse,
        Navbar,
        NavbarToggler,
        NavbarBrand,
        Nav,
        NavItem,
        NavLink,
        UncontrolledDropdown,
        DropdownToggle,
        DropdownMenu,
        DropdownItem
    } from 'sveltestrap';

    export let isLoggedOn = false;
    let servicesDropDown = {masterData:["Customers", "Products"], authentication:["Users"] }
    let servicesTopLevel = ["Promotions"]

    let isOpen = false;

    function handleUpdate(event) {
        isOpen = event.detail.isOpen;
    }
</script>

<Navbar color="light" light expand="md">
    <NavbarBrand href="/">goTemp</NavbarBrand>
    <NavbarToggler on:click={() => (isOpen = !isOpen)} />
    <Collapse {isOpen} navbar expand="md" on:update={handleUpdate}>
        <Nav class="ml-auto" navbar>
            {#each servicesTopLevel as service}
                <NavItem>
                    <NavLink href="https://bitbucket.org/Bolbeck/gotemp">{service}</NavLink>
                </NavItem>
            {/each}
            <UncontrolledDropdown nav inNavbar>
                <DropdownToggle nav caret>Services</DropdownToggle>
                <DropdownMenu right>
                    {#each servicesDropDown.masterData as service}
                        <DropdownItem>{service}</DropdownItem>
                    {/each}
                    <DropdownItem divider />
                    {#each servicesDropDown.authentication as service}
                        <DropdownItem>{service}</DropdownItem>
                    {/each}
                </DropdownMenu>
            </UncontrolledDropdown>
            <NavItem>
                {#if isLoggedOn}
                    <NavLink href="#components/">Logout</NavLink>
                {:else}
                    <NavLink href="#components/">Login</NavLink>
                {/if}
            </NavItem>
            <NavItem>
                <NavLink href="https://bitbucket.org/Bolbeck/gotemp"><i class="fab fa-bitbucket"></i></NavLink>
            </NavItem>
        </Nav>
    </Collapse>
</Navbar>