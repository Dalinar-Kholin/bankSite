import {useEffect, useState} from "react";
import Transfer from "./Transfer.tsx";
import useCheckCookie from "../customHook/useCheckCookie.ts";


interface Transaction {
    Sender: number;
    Reciver: number;
    Value: number;
}
export default function LoadTransferSent(){
//'https://127.0.0.1:8080/przelewy'
    const [data, setData] = useState<Transaction[]>([]);
    useCheckCookie()

    useEffect(() => {
        fetch('https://127.0.0.1:8080/api/transfer',{
            method: 'GET',
            credentials: 'include'
        })
            .then(response => response.json())
            .then(fetchedData => {
                if (fetchedData==null){
                    return
                }
                setData(fetchedData);
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });

        // Pusty array zależności oznacza, że efekt uruchomi się tylko raz przy montowaniu komponentu
    }, []);
    var i=0
    return(
        <>
            liczba znalezionych przelewów := {data.length}
            <div className={"nice"}>
                {
                    data === null?
                        "there is no Transfer" :
                        data.map(item =>
                            <Transfer key={i++} from={item.Sender} to={item.Reciver} amount={item.Value}/>)
                }
            </div>
        </>
    )
}