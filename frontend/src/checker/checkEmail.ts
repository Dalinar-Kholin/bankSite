export default function isValidEmail(email: string): boolean {
    const pattern: RegExp = /^[a-zA-Z0-9_+&*-]+(?:\.[a-zA-Z0-9_+&*-]+)*@(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$/;
    return pattern.test(email);
}