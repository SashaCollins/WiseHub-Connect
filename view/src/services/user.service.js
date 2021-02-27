import axios from 'axios';
import authHeader from "./auth-header";
import sha3 from 'crypto-js/sha3';

axios.defaults.xsrfCookieName = 'csrftoken'
axios.defaults.xsrfHeaderName = "X-CSRFTOKEN"

const HOST_API_URL = 'http://localhost:9010';
const DOCKER_API_URL = '/api';
const API_USER_URL = HOST_API_URL + '/user/';
const API_DATA_URL = HOST_API_URL + '/data/';


class UserService {

  fetchProfile(user) {
   return axios.get(API_USER_URL + 'profile', {
     headers: authHeader(),
     withCredentials: true,
     credentials: 'include'
   }
  ).then((response) => {
     return response;
   });
  }
  
  updateEmail(payload) {
    return axios.post(API_USER_URL + 'update/email', {
      option: 'email',
      new_email: payload.newEmail
    }, {
      headers: authHeader(),
      withCredentials: true,
      credentials: 'include'
    })
  }
  
  updateCredentials(payload) {
    return axios.post(API_USER_URL + 'update/credentials', {
      option: 'credentials',
      plugins: payload.plugins
    }, {
      headers: authHeader(),
      withCredentials: true,
      credentials: 'include'
    })
  }
  
  updatePassword(user) {
    return axios.post(API_USER_URL + 'update/password', {
      option: 'password',
      password: sha3(user.password).toString()
    }, {
      headers: authHeader(),
      withCredentials: true,
      credentials: 'include'
    })
  }
  
  deleteAccount(user) {
    return axios.get(API_USER_URL + 'delete', {
      headers: authHeader(),
      withCredentials: true,
      credentials: 'include'
    })
  }
  
  fetchData(payload) {
    return axios.post(API_DATA_URL + "all", {
          option: payload.option,
        }, {
          headers: authHeader(),
          withCredentials: true,
          credentials: 'include'
        }
    ).then((response) => {
      return response;
    });
  }
}

export default new UserService();
