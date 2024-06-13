import useCheckCookie from "../customHook/useCheckCookie.ts";
import {useState} from "react";

interface IResponse{
    link : string
    from : string
    to : string
    amount : number
}


export default function MakeTransfer(){
    useCheckCookie()
    const initial : IResponse =  {
        link : "",
        from : "",
        to : "",
        amount : 0,
    }
    const [respones, setRespones] = useState<IResponse>(initial)
    const [to, setTo] = useState<string>("")
    const [amount, setAmount] = useState<string>("")
    const [isOK, setIsOK] = useState<boolean>(false)
    var isResponseOK = false
    return(
        <>
            <form
                onSubmit={e=>
            {
                e.preventDefault()
                const request = {
                    reciver: to,
                    amount: +amount,
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
                            isResponseOK=false
                            return response.json()
                        }
                        isResponseOK= true
                        return response.json(); // Przetwarzanie odpowiedzi JSON tylko gdy response.ok
                    })
                    .then(data => {
                        console.log(isResponseOK)
                        if (isResponseOK){
                            console.log(data)
                            setRespones(data)
                            setIsOK(true)
                        }else{
                            setRespones(prevState => ({
                                ...prevState,
                                link: data.error
                            }))
                            setIsOK(false)
                        }
                    })
                    .catch(error => {
                        setRespones( prev => ({
                            ...prev,
                            link : "is not OK"
                        }))
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
                <p>{ isOK ? <a href={respones.link}>link do potwierdzenia transakcji</a> : respones.link ? respones.link : "nie udało się :(("}</p>
            </form>
        </>
    )
}