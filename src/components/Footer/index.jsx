import React from 'react';
import './assets/style.scss';
import {Button, Row, Container, Input} from 'reactstrap';

const Footer = props => {
    return(
        <footer className="footer fixed-bottom bg-light">
            <Container>
                <Row className="text-center py-3">
                    <Input type="text" className="form-control-lg ml-2 mr-4 w-75" placeholder="Create a task here"> </Input>
                    <Button type="submit" className="btn btn-success mr-2">Submit</Button>
                </Row>
            </Container>
        </footer>
    )
}

export default Footer