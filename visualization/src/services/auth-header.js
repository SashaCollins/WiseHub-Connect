export default function authHeader() {
    let token = sessionStorage.getItem('user');

    if (token) {
        return { Authorization: 'Bearer ' + token };
    } else {
        return {};
    }
}