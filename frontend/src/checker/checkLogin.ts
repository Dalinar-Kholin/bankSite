export default function isValidUsername(username: string): boolean {
    const regex = /^[a-zA-Z0-9]{5,20}$/;
    return regex.test(username);
}