import { useState } from "react"
import { Redirect } from "react-router-dom"


const Register = () => {
    const [username,setName] = useState()
    const [emailid,setEmail] = useState()
    const [password,setPassword] = useState()
    const [redirect,setRedirect] = useState(false)
    
    const submit = async e => {
        e.preventDefault()
        
        const response = await fetch('/api/register',{
            method: 'POST',
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify({
                username,
                emailid,
                password
            })
        });

        //const output = await response.json()

        if(response.status === 200){
            setRedirect(true)
        }
    }
    
    // eslint-disable-next-line
   if(redirect){
    return(
        <Redirect to="/login"></Redirect>
    )
   }

    return(
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please Register</h1>

            
            <input className="form-control" placeholder="Name" required onChange={e=>setName(e.target.value)}/>
        

            
            <input type="email" className="form-control" placeholder="Email Address" required onChange={b=>setEmail(b.target.value)}/>

            

            <input type="password" className="form-control" placeholder="Password" required onChange={e=>setPassword(e.target.value)}/>


            <button className="w-100 btn btn-lg btn-primary" type="submit" >register</button>

       </form> 
    )
}

export default Register