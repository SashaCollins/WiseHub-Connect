import axios from 'axios';
// import authHeader from "./auth-header";

const API_URL = 'http://localhost:9010/auth/';

class AuthService {
    login(user) {
        let hashed_pw = require('crypto').createHash('sha512')
            .update(user.password).digest('hex');
        return axios.post(API_URL + 'signin', {
            email: user.email,
            password: hashed_pw
        }).then((response) => {
            if (response.data) {
                sessionStorage.setItem('user', response.data.access);
            }
            console.log(sessionStorage.getItem('user'));
            console.log(response);
            user.password = '';
            return response;
        });
    }

    logout() {
        axios.get(API_URL + 'logout').then(() => {
            sessionStorage.clear();
            localStorage.clear();
        });
    }

    register(user) {
        let hashed_pw = require('crypto').createHash('sha512')
            .update(user.password).digest('hex');
        return axios.post(API_URL + 'signup', {
            name: user.name,
            password: hashed_pw,
            email: user.email
        }, {
            "content-type": "text/plain"
        });
    }
}

export default new AuthService();