import {useCallback, useEffect, useState} from "react";
import TransferComp from "./TransferComp.tsx";
import fetchWithSession from "../sessionFetch.ts";
export interface Itransfers{
    From: string,
    To: string,
    Title: string,
    Amount: number
}



interface IGetData{
    link: string
}
/*
interface ITransferData{
    receiver: string,
    sender: string,
    title: string,
    value: number,
}*/

/*
function switcher(data: any){
    let result =[];
    for (let i = 0;i<data.length;i++){
        let item : ITransferData ={
            receiver: data.To,
            sender: data.From,
            title: data.Title,
            value: data.Amount
        }
        result.push(item)
    }
    return result;
}
*/


async function getData({link}: IGetData){
    const response = await fetchWithSession(link,{
        credentials: 'include'
    });
    /*const responseData = await response.text()
    const data = eval(responseData);*/
    const data = await response.json()

    if (!response.ok) {
        throw new Error(`HTTP error! status: ${data.error}`);
    }
    return data;
}
interface IAcceptTransfers{
    link : string,
}

export default function AcceptTransfers({link}: IAcceptTransfers){
    const [transfers, setTransfers] = useState<Itransfers[]>([])
    const [loading, setLoading] = useState<boolean>(false)
    const [error, setError] = useState<string | null>(null)
    const fetchData = useCallback(
        async function fetchData() {
            setLoading(true);
            setError(null);
            try {
                const jsonData = await getData({link: link});
                setTransfers(jsonData);
            } catch (error: any) {
                setError(error.message);
            } finally {
                setLoading(false);
            }
        },
        [link]
    );

    useEffect(() => {
        fetchData();
    }, [fetchData]);

    if (loading){
        return <p key={"loading"}>Loading</p>
    }
    if (error) {
        return <p key={"error"}>Error: {error}</p>;
    }//javascript:document.getElementById('Kacper OsadowskiKacper Osadowski0').click()

    return (
        <>

            <h1> accept transfers</h1>
            {transfers ? transfers.map((t, index) =>
                    <TransferComp key={index} data={t}/>)
                :
                <></>
            }
        </>
    )
}