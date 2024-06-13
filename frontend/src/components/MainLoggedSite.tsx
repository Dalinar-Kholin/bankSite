import {useNavigate} from "react-router-dom";
import useCheckCookie from "../customHook/useCheckCookie.ts";

export default function MainLoggedSite() {
    const navigate = useNavigate();
    useCheckCookie()


    return (
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
            <p></p>
            {
                sessionStorage.getItem("isAdmin")==="true"
                    ?
                    <button onClick={() => {navigate("/acceptTransfers")}}>
                        accept Transfers
                    </button>
                    :
                    <></>}
        </>
    )
}