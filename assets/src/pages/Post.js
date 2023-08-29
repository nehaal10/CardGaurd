import { useState } from "react";


const Post = () => {

    const [creditcardno,setCredit] = useState()
    const [cvv,setCvv] = useState()
    const [expirydate,setExp] = useState()
    const [content,setContent] = useState()
    const [state,setStatus] = useState(false)

    const submit = async e => {
        e.preventDefault()
        const res = await fetch('/api/post',{
            method:'POST',
            headers:{'Content-Type':'application/json'},
            credentials:'include',
            body:JSON.stringify({
                creditcardno,
                cvv,
                expirydate
            })
        })

        await fetch('api/logout',{
            credentials:'include'
        })
        
        setContent(await res.json())
        setStatus(true)
    }


    if(state) {
        return (

            <h1>{content}</h1> 
        )
    }

    return(
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">ENTER DETAILS</h1>
            <input className="form-control" placeholder="number" required onChange={e=>setCredit(e.target.value)}/>
            <input className="form-control" placeholder="cvv" required onChange={e=>setCvv(e.target.value)}/>
            <input className="form-control" placeholder="exipry date" required onChange={e=>setExp(e.target.value)}/>
            <button className="w-100 btn btn-lg btn-primary" type="submit">submit</button>
        </form>
    )
}

export default Post;