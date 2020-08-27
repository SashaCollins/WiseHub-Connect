import UserService from '../services/user.service';

const userObject = sessionStorage.getItem('user');
const initialState = userObject
	? { user: userObject }
	: { user: null };

export const user = {
  namespaced: true,
  state: initialState,
  actions: {
	setUserState({commit}, user) {
		commit("setUserState", user);
	},
	fetchProfile({ commit }, user) {
	  return UserService.fetchProfile(user).then(onSuccess => {
	    if (onSuccess.data.success) {
	      user.plugins = onSuccess.data.plugins;
		  commit("fetchSuccess", user);
		}
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},
	updateEmail({ commit }, payload) {
	  return UserService.updateEmail(payload).then(onSuccess => {
	    let updatedUser = user.state.status.user.email = payload.newEmail;
		commit("updateSuccess",updatedUser);
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
		user = null;
		sessionStorage.clear();
		localStorage.clear();
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	}
  },
  mutations: {
	setUserState(state, user) {
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
	}
  },
}


