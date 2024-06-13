import './App.css'
import LoginPage from "./components/Login.tsx";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import MainLoggedSite from "./components/MainLoggedSite.tsx";
import RegisterNewUser from "./components/registerNewUser.tsx";
import Redirect from "./components/Redirect.tsx";
import ResetPasswd from "./components/resetPasswd.tsx";
import LoadTransferSent from "./components/LoadTransferSent.tsx";
import MakeTransfer from "./components/initNewMoneyTransfer.tsx";
import AcceptTransfers from "./components/acceptTransfers.tsx";

function App() {
    return (
        <>

        <Router>
            <Routes>
                <Route path="/register" element={<RegisterNewUser />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="/" element={<Redirect />} />
                <Route path="/main" element={<MainLoggedSite />} />
                <Route path="/resetPassword" element={<ResetPasswd />} />
                <Route path={"/getTransfer"} element={<LoadTransferSent/>}/>
                <Route path={"/makeTransfer"} element={<MakeTransfer/>}/>
                <Route path="/acceptTransfers" element={<AcceptTransfers link={"https://127.0.0.1:8080/acceptTransfer"}/>}/>{/* potwierdzanie przelewów*/}
            </Routes>
        </Router>
        </>
    );
}
export default App;

/*
* const request ={
                    "to" : "nice",
                    "value" : 100
                }

                axios.post("https://127.0.0.1:8080/transfers", request, {
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    withCredentials: true // Włącz to jeśli potrzebujesz obsługi ciasteczek/CORS
                }).then(r =>
                    setStatus(r.data)
                ).catch(
                error =>
                    setStatus(error)
                )*/