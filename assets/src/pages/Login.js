import { useState } from 'react' 
import { Redirect } from 'react-router-dom'


const Login = () => {

  const [username,setName] = useState()
  const [emailid,setEmail] = useState()
  const [password,setPassword] = useState()
  const [redirect,setRedirect] = useState(false)

  const submit =async e => {
    e.preventDefault()
   const res =  await fetch('/api/login',{
      method:'POST',
      credentials:'include',
      headers:{'Content-Type':'application/json'},
      body: JSON.stringify({
        username,
        emailid,
        password
      })
    })

    if(res.status === 200){
      setRedirect(true)
    }

  }

  if(redirect){
    return (
      <Redirect to="/post" />
    )
  }

  return (
       <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please Login</h1>
            
              <input class="form-control" placeholder="Name" required onChange={e=>setName(e.target.value)}/>
  
            
              <input type="email" class="form-control" placeholder="Email ID" required onChange={e=>setEmail(e.target.value)}/>
              
            
      
              <input type="password" class="form-control" placeholder="Password" required onChange={e=>setPassword(e.target.value)}/>
            
            
            <button className="w-100 btn btn-lg btn-primary" type="submit">submit</button>
       </form> 
    )
}

export default Login;