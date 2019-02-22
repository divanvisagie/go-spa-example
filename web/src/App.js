import React, { Component } from 'react'
import {Todos} from './Todos'
import logo from './logo.svg'
import './App.css'

export default class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          Todo App
        </header>

        <Todos></Todos>
      </div>
    )
  }
}

