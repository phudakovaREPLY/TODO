import React, {useState, useEffect} from 'react';
import {Button, Container, Input, Row, ListGroup, ListGroupItem} from 'reactstrap';
import './assets/style.scss';
import axios from "axios";

const Content = props => {

    const [task, setTask] = useState("")
    const [taskList, setTaskList] = useState(null)

    function getDataFromAPI(){
        axios.get("http://localhost:8080/api/v1/getall", {params: {
            done: false
            }}).then(res => {
            if(res.status===200){
                setTaskList(res.data)
            } else {
                alert("Something went wrong")
            }
        })
    }

    function handleSubmit() {
        axios.post("http://localhost:8080/api/v1/createtodo", {task}).then(res => {
            if(res.status===200){
                let tl = !taskList ? [] : taskList
                tl.push(task)
                setTaskList(tl)
                setTask("")
            } else {
                alert("Something went wrong")
            }
        })
    }

    function handleDeleteAll() {
        axios.delete("http://localhost:8080/api/v1/removealltodos",{params: {
                done: false
            }}).then(res => {
            if(res.status===200){
                getDataFromAPI()
            } else {
                alert("Something went wrong")
            }
        })
    }

    useEffect(() => {
        getDataFromAPI()
    }, []);

    const renderTaskList = []
    if(!taskList || taskList.length === 0) {
        renderTaskList.push(
            <ListGroupItem key="empty">OOPS... much empty </ListGroupItem>
        )
    }
    else {
        for(let i = 0; i < taskList.length; i++) {
            renderTaskList.push(
                <ListGroupItem key={`item-${i}`}>{taskList[i].task}</ListGroupItem>
            )
        }
    }


    return(
        <div className="content">
            <Container className= "mt-3 rounded border border-success">
                <ListGroup>
                    {renderTaskList}
                </ListGroup>
            </Container>

            <Container>
                <Row className="text-center py-3">
                    <Input type="text" name="task" value={task} onChange={e => setTask(e.target.value)} className="form-control-lg ml-2 mr-4 w-75" placeholder="Create a task here"> </Input>
                    <Button type="button" className="btn btn-success mr-2 mt-2" onClick={handleSubmit} >Submit</Button>
                    <Button type="button" className="btn btn-danger mr-2 mt-2" onClick={handleDeleteAll} >Delete All</Button>
                </Row>
            </Container>
        </div>
    )
}



export default Content