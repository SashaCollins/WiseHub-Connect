import axios from 'axios';

const API_URL = 'http://localhost:9020/admin/';

class UserService {

  fetchProfile(user) {
	return axios.post(API_URL + 'profile', {
		  email: user.email
		}
	).then((response) => {
	  return response;
	});
  }

  updatePlugins(payload) {
	window.console.error(payload);
	return axios.post(API_URL + 'update/plugins', {
	  option: 'plugins',
	  email: payload.email,
	  plugins: payload.plugins
	})
  }
  deletePlugins(payload) {
	return axios.post(API_URL + 'delete/plugins', {
	  option: 'plugins',
	  email: payload.email,
	  plugins: payload.plugins
	})
  }

}

export default new UserService();