import axios from 'axios';

const API_URL = 'http://localhost:9010/user/';

class UserService {

  fetchProfile(user) {
   return axios.post(API_URL + 'profile', {
     email: user.email
   }
  ).then((response) => {
     return response;
   });
  }

  updateEmail(payload) {
    return axios.post(API_URL + 'update/email', {
      option: 'email',
      email: payload.oldEmail,
      new_email: payload.newEmail
    })
  }

  // updatePlugins(payload) {
  //   window.console.error(payload);
  //   return axios.post(API_URL + 'update/plugins', {
  //     option: 'plugins',
  //     email: payload.email,
  //     plugins: payload.plugins
  //   })
  // }

  updatePassword(user) {
    let hashedPassword = require('crypto').createHash('sha512')
        .update(user.password).digest('hex');
    return axios.post(API_URL + 'update/password', {
      option: 'password',
      email: user.email,
      password: hashedPassword
    })
  }

  deleteAccount(user) {
    return axios.post(API_URL + 'delete', {
      email: user.email
    })
  }

  fetchRepos(user) {
    return axios.post(API_URL + 'repos', {
          email: user.email
        }
    ).then((response) => {
      return response;
    });
  }

  fetchCourses(user) {
    return axios.post(API_URL + 'all', {
          email: user.email
        }
    ).then((response) => {
      return response;
    });
  }

  fetchTeamRepos(payload) {
    return axios.post(API_URL + 'teams', {
          email: payload.user.email,
          organization: payload.organization
        }
    ).then((response) => {
      return response;
    });
  }
}

export default new UserService();