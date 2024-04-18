import {useState} from "react";
import { useNavigate } from 'react-router-dom';


export default function LoginBox(){
    const [userName, setUserName] = useState<string>("")
    const [pass, setPass] = useState<string>("")
    const [isSuccessfull, setIsSuccessfull] = useState<boolean>(true)
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
                    const request = {
                        login: userName,
                        pass: pass,
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
        </>

    )
}