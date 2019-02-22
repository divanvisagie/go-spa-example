import React, {useState} from 'react'
import './Todo.css'

export const Todo = ({handleAdd}) => {
    const [description, setDescription] = useState('')


    return (<form onSubmit={x => handleAdd({description})}>
        <input type="text" value={description} onChange={x => setDescription(x.target.value)}></input>
        <button disabled={description.length < 1}>Add</button>
    </form>)
}