import {useEffect, useState} from "react";
import Item from "./Item.tsx";
import useCheckCookie from "../customHook/useCheckCookie.ts";


interface Transaction {
    Sender: number;
    Reciver: number;
    Value: number;
}
export default function GetTransfers(){
//'https://127.0.0.1:8080/przelewy'
    const [data, setData] = useState<Transaction[]>([]);
    useCheckCookie()

    useEffect(() => {
        fetch('https://127.0.0.1:8080/przelewy',{
            method: 'GET',
            credentials: 'include'
        })
            .then(response => response.json())
            .then(fetchedData => {
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
                {data.map(item =>
                    <Item key={i++} from={item.Sender} to={item.Reciver} amount={666}/>)
                }
            </div>
        </>
    )
}