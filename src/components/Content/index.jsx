import React from 'react';
import {Container} from 'reactstrap';
import './assets/style.scss';

const Content = props => {
    return(
        <div className="content">
            <Container className= "mt-3 rounded border border-success">
                    <p className="mt-3 mx-3">Nothing here, much sad...
                        <br/><br/><br/><br/><br/><br/><br/><br/>
                        <br/><br/><br/><br/><br/><br/><br/><br/>
                        <br/><br/><br/><br/><br/><br/><br/><br/></p>
            </Container>
        </div>
    )
}



export default Content