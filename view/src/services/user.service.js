import axios from 'axios';

const API_USER_URL = '/api/user/';
const API_DATA_URL = '/api/data/';

class UserService {

  fetchProfile(user) {
   return axios.post(API_USER_URL + 'profile', {
     email: user.email
   }
  ).then((response) => {
     return response;
   });
  }
  updateEmail(payload) {
    return axios.post(API_USER_URL + 'update/email', {
      option: 'email',
      email: payload.oldEmail,
      new_email: payload.newEmail
    })
  }
  updateCredentials(payload) {
    return axios.post(API_USER_URL + 'update/credentials', {
      option: 'credentials',
      email: payload.email,
      plugins: payload.plugins
    })
  }
  updatePassword(user) {
    let hashedPassword = require('crypto').createHash('sha512')
        .update(user.password).digest('hex');
    return axios.post(API_USER_URL + 'update/password', {
      option: 'password',
      email: user.email,
      password: hashedPassword
    })
  }
  deleteAccount(user) {
    return axios.post(API_USER_URL + 'delete', {
      email: user.email
    })
  }
  fetchData(payload) {
    return axios.post(API_DATA_URL + "all", {
          option: payload.option,
          email: payload.user.email
        }
    ).then((response) => {
      return response;
    });
  }
}

export default new UserService();