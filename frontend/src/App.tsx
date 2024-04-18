import './App.css'
import LoginPage from "./components/LoginBox.tsx";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import MainSite from "./components/mainSite.tsx";
import RegisterBox from "./components/registerBox.tsx";
import Redirect from "./components/Redirect.tsx";
import ResetPasswd from "./components/resetPasswd.tsx";
import GetTransfers from "./components/getTransfers.tsx";
import MakeTransfer from "./components/newTransfer.tsx";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/register" element={<RegisterBox />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="/" element={<Redirect />} />
                <Route path="/main" element={<MainSite />} />
                <Route path="/resetPassword" element={<ResetPasswd />} />
                <Route path={"/getTransfer"} element={<GetTransfers/>}/>
                <Route path={"/makeTransfer"} element={<MakeTransfer/>}/>
            </Routes>
        </Router>
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