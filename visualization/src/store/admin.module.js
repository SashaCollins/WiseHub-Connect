import AdminService from '../services/admin.service';

const adminObject = sessionStorage.getItem('admin');
const initialState = adminObject
	? { admin: adminObject }
	: { admin: null };

export const admin = {
  namespaced: true,
  state: initialState,
  actions: {
	initAdmin({commit}, admin) {
	  commit("initAdmin", admin);
	},
	fetchProfile({ commit }, admin) {
	  return AdminService.fetchProfile(admin).then(onSuccess => {
		if (onSuccess.data.success) {
		  admin.plugins = onSuccess.data.plugins;
		  commit("fetchSuccess", admin);
		}
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},
	updatePlugins({ commit }, payload) {
	  return AdminService.updatePlugins(payload).then(onSuccess => {
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},
	deletePlugins({ commit }, payload) {
	  return AdminService.deletePlugins(payload).then(onSuccess => {
		commit("deleteSuccess");
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},

  },
  mutations: {
	initUser(state, admin) {
	  state.admin = admin;
	},
	fetchSuccess(state, admin) {
	  state.admin = admin;
	},
	updateSuccess(state, admin) {
	  state.admin = admin;
	},
	deleteSuccess(state) {
	  state.user = null;
	}
  },
}


