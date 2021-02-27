import axios from 'axios';
import { secure } from './encryption.service';
import sha3 from 'crypto-js/sha3';
import refreshHeader from "@/services/refresh-header";

const HOST_API_URL = 'http://localhost:9010';
const DOCKER_API_URL = '/api';
const AUTH_API_URL = HOST_API_URL + '/api/auth/';


class AuthService {
    login(user) {
        return axios.post(AUTH_API_URL + 'signin', {
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
        return axios.post(AUTH_API_URL + 'signup', {
            name: user.name,
            password: sha3(user.password).toString(),
            email: user.email
        },
        );
    }

    refresh() {
        return axios.get(AUTH_API_URL + 'refresh', {
            headers: refreshHeader(),
            withCredentials: true,
            credentials: 'include'
        });
    }
    
    validate_token(token) {
        return axios.get(AUTH_API_URL + 'validate/' + token);
    }
}

export default new AuthService();
