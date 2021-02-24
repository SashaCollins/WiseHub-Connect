import axios from 'axios';
import authHeader from "./auth-header";
import sha3 from 'crypto-js/sha3';


const API_URL = '/api';
const API_USER_URL = API_URL + '/user/';
const API_DATA_URL = API_URL + '/data/';


class UserService {

  fetchProfile(user) {
   return axios.post(API_USER_URL + 'profile', {
     email: user.email
   }, { headers: authHeader() }
  ).then((response) => {
     return response;
   });
  }
  
  updateEmail(payload) {
    return axios.post(API_USER_URL + 'update/email', {
      option: 'email',
      email: payload.oldEmail,
      new_email: payload.newEmail
    }, { headers: authHeader() })
  }
  
  updateCredentials(payload) {
    return axios.post(API_USER_URL + 'update/credentials', {
      option: 'credentials',
      email: payload.email,
      plugins: payload.plugins
    }, { headers: authHeader() })
  }
  
  updatePassword(user) {
    //let hashedPassword = require('crypto').createHash('sha512').update(user.password).digest('hex');
    return axios.post(API_USER_URL + 'update/password', {
      option: 'password',
      email: user.email,
      password: sha3(user.password).toString()
    }, { headers: authHeader() })
  }
  
  deleteAccount(user) {
    return axios.post(API_USER_URL + 'delete', {
      email: user.email
    }, { headers: authHeader() })
  }
  
  fetchData(payload) {
    return axios.post(API_DATA_URL + "all", {
          option: payload.option,
          email: payload.user.email
        }, { headers: authHeader() }
    ).then((response) => {
      return response;
    });
  }
}

export default new UserService();
