import {useEffect, useState} from "react";
import { useNavigate } from 'react-router-dom';
import GoogleLogin, {GoogleLoginResponse, GoogleLoginResponseOffline} from "react-google-login";
import {gapi} from "gapi-script";
interface GoogleSignInComponentProps {
    loginSuccess: (response: GoogleLoginResponse | GoogleLoginResponseOffline) => void;
}


export default function Login(){
    const [userName, setUserName] = useState<string>("")
    const [pass, setPass] = useState<string>("")
    const [isSuccessfull, setIsSuccessfull] = useState<boolean>(true)
    const navigate = useNavigate();


    useEffect(()=>{
        function start(){
            gapi.client.init({
                clientId:'668295700477-b47dpenr2fqno4ue5aj7s3tiojbh59mj.apps.googleusercontent.com',
                scope:""
            })
        }
        gapi.load('client:auth2',start)
    })

    const log : GoogleSignInComponentProps = {loginSuccess: (response: GoogleLoginResponse | GoogleLoginResponseOffline) => {
            if ('tokenId' in response) {

                const onlineResponse = response as GoogleLoginResponse;
                // Wyślij token ID do serwera
                fetch('https://127.0.0.1:8080/google', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        "Session-Id": ""+ 12
                    },
                    body: JSON.stringify( onlineResponse.profileObj),
                    credentials: 'include'
                })
                    .then(response => {
                        return response.json()
                    })
                    .then(data  => {
                            sessionStorage.setItem("Session-Id", data.session)
                            sessionStorage.setItem("isAdmin", data.isAdmin)
                            sessionStorage.setItem("jwt", data.jwt)
                            navigate("/main")
                        }

                    )
                    .catch(error => console.error('Error:', error));
            } else {
                console.log('Login received offline access code');
            }
        }
    }


    return(
        <>
            <h1>
                login box
            </h1>
            <form
                style={{borderRadius: '15px'}}
                onSubmit={(e) => {
                    e.preventDefault();
 /*                   const request = {
                        login: isValidUsername(userName)? userName : "",
                        pass: isValidPassword(pass) ? pass : "",
                    }*/
                    const request = { // zmienione na czas sql injection
                        login:  userName ,
                        pass:  pass ,
                    }
                    fetch('https://127.0.0.1:8080/login', {
                        method: 'POST',
                        body: JSON.stringify(request),
                        credentials: 'include'
                    })
                        .then(response => {
                            if (!response.ok) {
                                setUserName("");
                                setPass("");
                                throw new Error('Network response was not ok.');
                            }
                            return response.json(); // Przetwarzanie odpowiedzi JSON tylko gdy response.ok
                        })
                        .then(data => {
                            console.log(data);
                            navigate("/main"); // Przekierowanie następuje tylko tutaj

                        })
                        .catch(error => {
                            setUserName(""); // Czyszczenie danych użytkownika przy pomyślnym logowaniu
                            setPass("");
                            setIsSuccessfull(false)
                            console.error('There was a problem with your fetch operation:', error);
                            // Obsługa błędów logowania
                        });

                }}>
                <p>login</p>
                <input value={userName} onChange={(e) => setUserName(e.target.value)}/>
                <p>password</p>
                <input value={pass} onChange={(e) => setPass(e.target.value)}/>
                <p style={{color: !isSuccessfull ? "red" : "black"}}>{isSuccessfull ? "" : "bad data"}</p>
                <button type="submit">{"sign in"}</button>
                <button onClick={()=>{
                    navigate("/register")
                }}>{"register"}</button>
                <button onClick={()=>{
                    navigate("/resetPassword")
                }}>forgot password</button>
            </form>

                <GoogleLogin
                    clientId={"668295700477-b47dpenr2fqno4ue5aj7s3tiojbh59mj.apps.googleusercontent.com"}
                    buttonText='Google'
                    onSuccess={log.loginSuccess}
                    responseType='code,token'
                    onFailure={(response: GoogleLoginResponse | GoogleLoginResponseOffline) => {
                        console.log(response)
                    }
                }
                />

        </>

    )
}