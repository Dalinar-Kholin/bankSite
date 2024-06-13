import {useState} from "react";
import { useNavigate } from 'react-router-dom';
import isValidUsername from "../checker/checkLogin.ts";
import isValidPassword from "../checker/checkPass.ts";
import isValidEmail from "../checker/checkEmail.ts";


export default function RegisterNewUser(){
    const [userName, setUserName] = useState<string>("")
    const [pass, setPass] = useState<string>("")
    const [email, setEmail] = useState<string>("")
    const [repPass, setRepPass] = useState<string>("")
    const navigate = useNavigate();
    return(
        <>
            <h1>
            login box
            </h1>
            <form
                style={{borderRadius: '15px'}}
                onSubmit={(e) => {
                    e.preventDefault();
                    navigate("/main")
                    const request = {
                        login : isValidUsername(userName) ? userName : "",
                        pass : isValidPassword(pass)? pass : "",
                        email : isValidEmail(email) ? email : ""
                    }

                    fetch('https://127.0.0.1:8080/api/registerNewUser', {
                        method : 'Post',
                        body : JSON.stringify(request)
                    })
                        .then(response => {
                            if (!response.ok) {
                                throw new Error('Network response was not ok');
                            }
                            return response.json();
                        })
                        .then(data => console.log(data))
                        .catch(error => console.error('There was a problem with your fetch operation:', error));
                    navigate("/login")
                    setUserName("")
                    setEmail("")
                    setPass("")
                    setRepPass("")
                }}>
                <p>login</p>
                <input value={userName} onChange={(e) => setUserName(e.target.value)}/>
                <p>email</p>
                <input value={email} onChange={(e) => setEmail(e.target.value)}/>
                <p>password</p>
                <input value={pass} onChange={(e) => setPass(e.target.value)}/>
                <p>repeat pass</p>
                <input value={repPass} onChange={(e) => setRepPass(e.target.value)}/>
                <p></p>
                <button type="submit">{"add Account"}</button>
            </form>
        </>

    )
}