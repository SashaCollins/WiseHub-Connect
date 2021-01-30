import axios from 'axios';

const API_USER_URL = 'http://localhost:9010/user/';
const API_DATA_URL = 'http://localhost:9010/data/';
const API_ADMIN_URL = 'http://localhost:9010/admin/';

class UserService {

  updatePlugins(payload) {
    window.console.error(payload);
    return axios.post(API_ADMIN_URL + 'update/plugins', {
      option: 'plugins',
      email: payload.email,
      plugins: payload.plugins
    })
  }
  deletePlugins(payload) {
    return axios.post(API_ADMIN_URL + 'delete/plugins', {
      option: 'plugins',
      email: payload.email,
      plugins: payload.plugins
    })
  }

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
    window.console.error(payload);
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

  fetchRepos(user) {
    return axios.post(API_USER_URL + 'repos', {
          email: user.email
        }
    ).then((response) => {
      return response;
    });
  }

  fetchCourses(user) {
    return axios.post(API_DATA_URL + "all", {
          option: "general",
          email: user.email
        }
    ).then((response) => {
      return response;
    });
  }

  fetchTeamRepos(payload) {
    return axios.post(API_USER_URL + 'teams', {
          email: payload.user.email,
          organization: payload.organization
        }
    ).then((response) => {
      return response;
    });
  }
}

export default new UserService();