import axios from 'axios';
import { secure } from './encryption.service';
import sha3 from 'crypto-js/sha3';


const API_URL = '/api/auth/';


class AuthService {
    login(user) {
        return axios.post(API_URL + 'signin', {
            email: user.email,
            password: sha3(user.password).toString()
        }).then((response) => {
            //set user to loggedIn
            if (response.data) {
                if (response.data.success) {
                    user.password = "";
                    secure.set('user', JSON.stringify(user));
                    secure.set('token', response.data.token);
                    secure.set('loggedIn', response.data.success);
                }
            }
            return response;
        });
    }

    logout() {
        secure.removeAll();
        sessionStorage.clear();
        localStorage.clear();
    }

    register(user) {
        return axios.post(API_URL + 'signup', {
            name: user.name,
            password: sha3(user.password).toString(),
            email: user.email
        },
        );
    }
    
    validate_token(token) {
        return axios.get(API_URL + 'validate/' + token);
    }
}

export default new AuthService();
