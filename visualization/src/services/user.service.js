import axios from 'axios';

const API_URL = 'http://localhost:9010/user/';

class UserService {
  fetchProfile(user) {
   return axios.post(API_URL + 'profile', {
     email: user.email
   }
  );
  }

  updateEmail(payload) {
    return axios.post(API_URL + 'update/email', {
      old_email: payload.oldEmail,
      new_email: payload.newEmail
    })
  }

  updatePassword(user) {
    let hashedPassword = require('crypto').createHash('sha512')
        .update(user.password).digest('hex');
    return axios.post(API_URL + 'update/password', {
      email: user.email,
      password: hashedPassword
    })
  }

  deleteAccount(user) {
    return axios.post(API_URL + 'delete', {
      email: user.email
    })
  }
}

export default new UserService();