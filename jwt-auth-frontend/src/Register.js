import React , {useState} from 'react'

import axios from 'axios'

const Register = () => {

    const [username , setUsername] =  useState('')
    const [password , setPassword] =  useState('')

    const handleRegister = async (e) =>{
        e.preventDefault();

        try{
            await axios.post('/register', {username,password})
            alert('Registration successful')

        }
        catch(error){
            console.error('Registration error', error);
            alert('Registration failed')
        }
    }

    return (
        <div>
            <h2>
                Register
            </h2>

            <form onSubmit={handleRegister}>
                <input type='text' placeholder='Username'  value={username} onChange={(e)=> setUsername(e.target.value)} required/>

                <input type='password' placeholder='password'  value={password}
                onChange={(e)=> setPassword(e.target.value)} required />

                <button type='submit'>Register</button>
            </form>
        </div>
    )
   
}

export default Register;