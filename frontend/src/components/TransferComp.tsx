import {Itransfers} from "./acceptTransfers.tsx";
import styled from "@emotion/styled";
import {Button} from "@mui/material";
interface ITransferComp{
    data : Itransfers
}


const StyledDiv = styled.div`
    background-color: #1a1a1a;
    padding: 10px 30px;
    border-radius: 20px;
    margin: 10px;
    display: flex;
    gap: 10px;
`

export default function TransferComp({data}:ITransferComp ){
    return(
        <>
            <StyledDiv>
                <p>{data.From} {"-- >"} {data.To} {data.Amount} {"  --->  "} {data.Title}</p>
                <Button
                    id={data.From+ data.To+data.Amount}
                    key={"accept"}
                    onClick={()=>{
                    alert("sent")
                }} variant={"text"} color={"success"}>accept</Button>
            </StyledDiv>
        </>
    )
}