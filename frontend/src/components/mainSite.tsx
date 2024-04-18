import {useNavigate} from "react-router-dom";
import useCheckCookie from "../customHook/useCheckCookie.ts";

export default function MainSite(){
    const navigate = useNavigate();
    useCheckCookie()

    return(
        <>
            <button onClick={() => {
                navigate("/makeTransfer")
            }}>make tranfer
            </button>
            <p></p>
            <button onClick={() => {
                navigate("/getTransfer")
            }}>get transfers
            </button>
        </>
    )
}