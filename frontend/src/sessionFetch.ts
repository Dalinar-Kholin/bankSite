export default async function fetchWithSession(url: string, options: RequestInit = {}): Promise<Response> {
    const sessionID = sessionStorage.getItem('Session-Id');
    if (!sessionID) {
        throw new Error('No active session');
    }
    const jwt = sessionStorage.getItem('jwt');
    if (!jwt) {
        throw new Error('No active session');
    }

    // Tworzenie obiektu Headers, jeśli nie istnieje w options
    const headers = new Headers(options.headers || {});
    headers.set('Session-Id', sessionID);
    headers.set('Content-Type', 'application/json');
    headers.set('Authorization', 'Bearer ' + jwt)
    const newOptions: RequestInit = {
        ...options,
        headers: headers
    };


    return fetch(url, newOptions);
}
