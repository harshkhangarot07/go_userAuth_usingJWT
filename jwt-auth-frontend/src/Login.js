import React from "react";
import { useState } from "react";
import axios from "axios";

const  Login = () => {
    const[username , setUsername] = useState('')
    const[password , setPassword] = useState('')

    const handleLogin = async(e) =>{
        e.preventDefault();

        try{
            const response = await axios.post('/login' , {username , password})
            localStorage.setItem('authToken', response.data.token)
            alert('Login successful')
        }
        catch(error){
            console.error('Login error', error)
            alert('Login failed')
        }
    }

    return (
        <div>
            <h2>
                login
            </h2>

            <form onSubmit={handleLogin}>
                <input type="text" placeholder="username"  value={username} onChange={(e)=> setUsername(e.target.value)} required />
                <input type="password" placeholder="password"  value={password} onChange={(e)=> setPassword(e.target.value)} required />

                <button type="submit">login</button>
            </form>
        </div>
    )
}

export default Login;