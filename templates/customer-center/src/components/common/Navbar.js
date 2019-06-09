import React from 'react';
import {
    Navbar,
    NavbarBrand,
    NavbarToggler,
    Collapse,
    Nav,
    UncontrolledDropdown,
    DropdownToggle,
    DropdownMenu,
    DropdownItem,
    Form,
    FormGroup,
    Input,
    InputGroup
} from 'reactstrap';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function Header(props) {

    const [isOpen, setIsOpen] = React.useState(false);

    function toggleNavbar() {
        setIsOpen(!isOpen);
    }

    return (
        <Navbar color="primary" dark expand="md">
            <NavbarBrand href="/" className="text-white navbar-heading">
                Cartelize
            </NavbarBrand>
            <NavbarToggler onClick={toggleNavbar}></NavbarToggler>
            <Collapse isOpen={isOpen} navbar>
                <Nav className="ml-auto">
                    <Form inline>
                        <FormGroup className="mb-2 mr-sm-2 mb-sm-0">
                            <InputGroup>
                                <Input type="text" placeholder="Search" className="border-0 nav-search text-white bg-accent"></Input>
                                <div className="input-group-append">
                                    <span className="input-group-text nav-search-icon bg-accent border-0">
                                        <FontAwesomeIcon color="white" icon="search"></FontAwesomeIcon>
                                    </span>
                                </div>
                            </InputGroup>
                        </FormGroup>
                    </Form>
                    <UncontrolledDropdown nav inNavbar>
                        <DropdownToggle nav>
                            <FontAwesomeIcon color="white" size="1x" icon="shopping-cart" />
                        </DropdownToggle>
                        <DropdownMenu right>
                            <DropdownItem>
                                Option 1
                            </DropdownItem>
                            <DropdownItem>
                                Option 2
                            </DropdownItem>
                        </DropdownMenu>
                    </UncontrolledDropdown>
                    <UncontrolledDropdown nav inNavbar>
                        <DropdownToggle nav>
                            <FontAwesomeIcon color="white" size="1x" icon="user" />
                        </DropdownToggle>
                        <DropdownMenu right>
                            <DropdownItem href="/login">
                                SignIn / Register
                            </DropdownItem>
                        </DropdownMenu>
                    </UncontrolledDropdown>
                </Nav>
            </Collapse>
        </Navbar>
    )
}

export default Header;
