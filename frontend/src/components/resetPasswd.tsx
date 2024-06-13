import {useState} from "react";
import axios from "axios";
import isValidEmail from "../checker/checkEmail.ts";
import isValidUsername from "../checker/checkLogin.ts";

export default function ResetPasswd(){
    const [email, setEmail] = useState<string>("")
    const [login, setLogin] = useState<string>("")
    const [status, setStatus] = useState<string>("")
    return(
        <>
            <form
                onSubmit={(e) => {
                    axios.post("https://127.0.0.1:8080/api/resetPassword",{
                        login: isValidUsername(login)? login : "",
                        email: isValidEmail(email) ? email : "",
                    }).then(
                        (r)=>{
                            console.log(r.data)
                            if(r.data.isOK){
                                setStatus("link wysłany")
                            }else {
                                setStatus("bad data")
                            }
                        }
                    ).catch(()=>{
                        setStatus("bad data")
                    })
                    e.preventDefault()
                }}>
                <input
                    placeholder={"email"}
                    value={email} onChange={(e) => {
                    setEmail(e.target.value)
                }}></input>
                <p></p>
                <input
                    placeholder={"login"}
                    value={login} onChange={(e) => {
                    setLogin(e.target.value)
                }}></input>
                <p></p>
                <button type="submit">reset password</button>
                <p>{status}</p>
            </form>
        </>
    )
}