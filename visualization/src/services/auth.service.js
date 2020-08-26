import axios from 'axios';

const API_URL = 'http://localhost:9010/auth/';

class AuthService {
    login(user) {
        let hashedPassword = require('crypto').createHash('sha512')
            .update(user.password).digest('hex');
        return axios.post(API_URL + 'signin', {
            email: user.email,
            password: hashedPassword
        }).then((response) => {
            //set user to loggedIn
            console.log(response.data.Success)
            if (response.data) {
                if (response.data.Success) {
                    user.password = "";
                    sessionStorage.setItem('loggedIn', response.data.Success);
                    // browser session storage for user module
                    sessionStorage.setItem('user', JSON.stringify(user));
                }
            }
            console.log(sessionStorage.getItem('loggedIn'));
            console.log(response);
            return response;
        });
    }

    logout() {
        sessionStorage.clear();
        localStorage.clear();
    }

    register(user) {
        let hashedPassword = require('crypto').createHash('sha512')
            .update(user.password).digest('hex');
        return axios.post(API_URL + 'signup', {
            name: user.name,
            password: hashedPassword,
            email: user.email
        },
        );
    }
}

export default new AuthService();