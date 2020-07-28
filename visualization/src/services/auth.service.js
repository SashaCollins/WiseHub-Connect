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
            //set user to loggedIn
            console.log(response.data.Success)
            if (response.data) {
                if (response.data.Success) {
                    sessionStorage.setItem('loggedIn', response.data.Success);
                }
            }
            console.log(sessionStorage.getItem('loggedIn'));
            console.log(response);
            user.password = '';
            return response;
        });
    }

    logout() {
        sessionStorage.clear();
        localStorage.clear();
    }

    register(user) {
        let hashed_pw = require('crypto').createHash('sha512')
            .update(user.password).digest('hex');
        return axios.post(API_URL + 'signup', {
            name: user.name,
            password: hashed_pw,
            email: user.email
        },
        );
    }
}

export default new AuthService();