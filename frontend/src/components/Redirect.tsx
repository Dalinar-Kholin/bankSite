import {useNavigate} from "react-router-dom";
import { useEffect } from 'react';


export default function Redirect() {
    const navigate = useNavigate();

    useEffect(() => {
        navigate('/login');
    }, [navigate]);

    return null;  // lub może zwracać jakiś element ładowania, np. <div>Loading...</div>
}