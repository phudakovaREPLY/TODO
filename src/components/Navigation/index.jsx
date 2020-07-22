import React from 'react';
import './assets/style.scss';
import {Nav, NavbarBrand, Navbar, NavItem} from 'reactstrap';
import logo from './assets/replyLogo.png';

const Navigation = props => {
    return(
        <div className='navigation'>
            <Navbar className="navbar navbar-light bg-light">
                <Nav>
                    <NavItem>
                        <a href="/">
                            <img src={logo} height="33" width="33" alt="" /></a>
                    </NavItem>
                    <NavbarBrand>
                        TODO
                    </NavbarBrand>
                </Nav>
            </Navbar>
        </div>
    )
}



export default Navigation