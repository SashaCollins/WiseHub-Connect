import UserService from '../services/user.service';

const userObject = sessionStorage.getItem('user');
const initialState = userObject ? {status: {user: userObject}} : {status: {user: null}};

export const user = {
  namespaced: true,
  state: initialState,
  actions: {
	setUserState({commit}, user) {
		commit("setUserState", user)
	},
	fetchProfile({ commit }, user) {
	  return UserService.fetchProfile(user).then(response => {
		commit("fetchSuccess", user)
		return Promise.resolve(response);
	  }, error => {
		return Promise.reject(error);
	  });
	},
	updateEmail({ commit }, user) {
	  return UserService.updateEmail(user).then(response => {
		commit("updateSuccess", user)
		return Promise.resolve(response);
	  }, error => {
		return Promise.reject(error);
	  });
	},
	updatePassword({ commit }, user) {
	  return UserService.updatePassword(user).then(response => {
		user.password = '';
		commit("updateSuccess", user)
		return Promise.resolve(response);
	  }, error => {
		return Promise.reject(error);
	  });
	},
	deleteAccount({ commit }, user) {
	  return UserService.delete(user).then(response => {
		commit("deleteSuccess")
		user.password = '';
		sessionStorage.clear();
		localStorage.clear();
		return Promise.resolve(response);
	  }, error => {
		return Promise.reject(error);
	  });
	}
  },
  mutations: {
	setUserState(state, user) {
	  state.status.user = user;
	},
	fetchSuccess(state, user) {
	  state.status.user = user;
	},
	updateSuccess(state, user) {
	  state.status.user = user;
	},
	deleteSuccess(state) {
	  state.status.user = null;
	}
  },
}


