import useCheckCookie from "../customHook/useCheckCookie.ts";
import {useState} from "react";

export default function MakeTransfer(){
    useCheckCookie()
    const [to, setTo] = useState<string>("")
    const [amount, setAmount] = useState<string>("")
    const [status, setStatus] = useState<string>("")
    return(
        <>
            <form
                onSubmit={e=>
            {
                e.preventDefault()
                const request = {
                    reciver: to,
                    amount: amount,
                }

                fetch('https://127.0.0.1:8080/acceptTransaction', {
                    method: 'POST',
                    body: JSON.stringify(request),
                    credentials: 'include'
                })
                    .then(response => {
                        if (!response.ok) {
                            setTo("");
                            setAmount("");
                            throw new Error('Network response was not ok.');
                        }
                        return response.json(); // Przetwarzanie odpowiedzi JSON tylko gdy response.ok
                    })
                    .then(data => {
                        console.log(data);
                    })
                    .catch(error => {
                        setStatus("nie usaloi sie")
                        console.error('There was a problem with your fetch operation:', error);
                        // Obsługa błędów logowania
                    });

            }}
            >
                <p>to</p>
                <input value={to} onChange={(e) => setTo(e.target.value)}/>
                <p>amount</p>
                <input value={amount} onChange={(e) => setAmount(e.target.value)}/>
                <p></p>
                <button type="submit">{"send"}</button>
                <p>{status}</p>
            </form>
        </>
    )
}