import React, {Component, Fragment} from 'react'
import {Todo} from './Todo'

import './Todos.css'

export class Todos extends Component {
    constructor() {
        super()

        this.state = {
            todos: [],
            waiting: false
        }

        this.handleAdd = this.handleAdd.bind(this)
    }

    componentDidMount() {
        this.setState({waiting: true})
        fetch('/todo')
            .then(response => {
                this.setState({waiting: true})
                if (response.status == 200) {
                    return response.json()
                }
            })
            .then(todos => {
                this.setState({todos})
            })
            .finally(() =>{
                this.setState({waiting: false})
            })
    }

    handleAdd(todo) {
        if (!todo.description) return
        this.setState({waiting: true})
        fetch('/todo', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(todo)
        }) 
        .then(response => {
            this.setState({waiting: true})
            if (response.status == 200) {
                return response.json()
            }
        })
        .then(todos => {
            this.setState({todos})
        })
        .finally(() =>{
            this.setState({waiting: false})
        })
    }

    render() {

        const renderInput = () => {
            const {waiting} = this.state
            if (!waiting) {
                return (
                    <Todo handleAdd={this.handleAdd}></Todo>
                )
            }
            return (<div>Waiting...</div>)
        }

        return (
            <Fragment>
                <ul className="Todos">
                    {this.state.todos.map(x => (<li>
                        {x.description}
                    </li>))}
                </ul>
                {renderInput()}
            </Fragment>
        )
    }
}