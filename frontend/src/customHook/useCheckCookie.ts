import { useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const useCheckCookie = () => {
    const navigate = useNavigate();

    useEffect(() => {
        const checkCookie = async () => {
            try {
                const response = await axios.get('https://127.0.0.1:8080/api/checkCookie', {
                    withCredentials: true
                });
                if (!response.data.isValid) {
                    navigate('/login');
                }
            } catch (error) {
                console.error('Error checking cookie:', error);
                navigate('/login');
            }
        };

        checkCookie().then(()=>
                console.log("logged")
        );
    }, [navigate]);
};

export default useCheckCookie;