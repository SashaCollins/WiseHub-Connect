import UserService from '../services/user.service';
import Plugin from '../model/plugins';
import { secure } from '@/services/encryption.service'

const userObject = JSON.parse(sessionStorage.getItem('user'));
const initialState = userObject
	? { user: userObject }
	: { user: null };

export const user = {
  namespaced: true,
  state: initialState,
  actions: {
	initUser({commit}, user) {
		commit("initUser", user);
	},
	fetchProfile({ commit }, user) {
	  return UserService.fetchProfile(user).then(onSuccess => {
	    if (onSuccess.data.success) {
		  commit("fetchSuccess", user);
		}
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},
	updateCredentials({ commit }, payload) {
	  return UserService.updateCredentials(payload).then(onSuccess => {
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},
	updateEmail({ commit }, payload) {
	  return UserService.updateEmail(payload).then(
	  	onSuccess => {
	  	  user.state.user.email = payload.newEmail;
	  	  commit("updateSuccess", user);
	  	  return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},
	updatePassword({ commit }, user) {
	  return UserService.updatePassword(user).then(onSuccess => {
		user.password = '';
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},
	deleteAccount({ commit }, user) {
	  return UserService.delete(user).then(onSuccess => {
		commit("deleteSuccess");
		secure.removeAll();
		sessionStorage.clear();
		localStorage.clear();
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},
	fetchData({ commit }, payload) {
	  return UserService.fetchData(payload).then(onSuccess => {
		if (onSuccess.data.success) {
			if (onSuccess.data.pluginData) {
				payload.user.plugins = [];
				for (const [key, value] of Object.entries(onSuccess.data.pluginData)) {
					if (value) {
						let plugin = new Plugin(key, JSON.parse(value));
						payload.user.plugins.push(plugin);
					}
				}
			}
		  commit('fetchSuccess', payload.user);
		}
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},
  },
  mutations: {
	initUser(state, user) {
	  state.user = user;
	},
	fetchSuccess(state, user) {
	  state.user = user;
	},
	updateSuccess(state, user) {
	  state.user = user;
	},
	deleteSuccess(state) {
	  state.user = null;
	},
	deletePluginsSuccess(state, user) {
	  state.user = user;
	}
  },
}


