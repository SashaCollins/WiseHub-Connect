import axios from 'axios';
// import authHeader from "./auth-header";

const API_URL = 'http://localhost:10020/api/auth/';

class AuthService {
    login(user) {
        let hashed_pw = require('crypto').createHash('sha512').update(process.env.SALT_ONE + user.password + process.env.SALT_TWO).digest('hex');
        return axios.post(API_URL + 'signin', {
            email: user.email,
            password: hashed_pw
        }).then((response) => {
            if (response.data) {
                sessionStorage.setItem('user', response.data.access);
                sessionStorage.setItem('access-expire', response.data.expire)
            }
            console.log(sessionStorage.getItem('user'));
            console.log(response);
            user.password = '';
            return response;
        });
    }

    refresh() {
        return axios.get(API_URL + 'refresh').then((response) => {
            if (response.data.access) {
                sessionStorage.setItem('user', response.data.access);
                sessionStorage.setItem('access-expire', response.data.expire);
            }
            console.log(sessionStorage.getItem('user'));
            console.log(response);
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
        let hashed_pw = require('crypto').createHash('sha512').update(process.env.SALT_ONE + user.password + process.env.SALT_TWO).digest('hex');
        return axios.post(API_URL + 'signup', {
            name: user.username,
            email: user.email,
            password: hashed_pw
        });
    }
}

export default new AuthService();