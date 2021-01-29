import AdminService from '../services/admin.service';

const adminObject = sessionStorage.getItem('admin');
const initialState = adminObject
	? { user: adminObject }
	: { user: null };

export const admin = {
  namespaced: true,
  state: initialState,
  actions: {
	initUser({commit}, admin) {
	  commit("initUser", admin);
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
	updateEmail({ commit }, payload) {
	  return AdminService.updateEmail(payload).then(
		  onSuccess => {
			admin.state.user.email = payload.newEmail;
			commit("updateSuccess", admin);
			return Promise.resolve(onSuccess);
		  }, onFailure => {
			return Promise.reject(onFailure);
		  });
	},
	updatePassword({ commit }, admin) {
	  return AdminService.updatePassword(admin).then(onSuccess => {
		admin.password = '';
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},
	deleteAccount({ commit }, admin) {
	  return AdminService.delete(admin).then(onSuccess => {
		commit("deleteSuccess");
		sessionStorage.clear();
		localStorage.clear();
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		return Promise.reject(onFailure);
	  });
	},
	fetchRepos({ commit }, admin) {
	  return AdminService.fetchRepos(admin).then(onSuccess => {
		console.log(onSuccess)
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		console.log(onFailure)
		return Promise.reject(onFailure);
	  });
	},
	fetchCourses({ commit }, admin) {
	  return AdminService.fetchCourses(admin).then(onSuccess => {
		console.log(onSuccess)
		user.courses = [];
		if (onSuccess.data.success) {
		  for (const [key, value] of Object.entries(onSuccess.data.courses)) {
			console.log(`${key}: ${value}`);
			for (let i = 0; i < value.length; i++) {
			  user.courses.push(value[i]);
			}
		  }
		  commit('fetchSuccess', admin);
		}
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		console.log(onFailure)
		commit('fetchFailure');
		return Promise.reject(onFailure);
	  });
	},
	fetchTeams({ commit }, admin) {
	  return AdminService.fetchTeamRepos(admin).then(onSuccess => {
		console.log(onSuccess)
		return Promise.resolve(onSuccess);
	  }, onFailure => {
		console.log(onFailure)
		return Promise.reject(onFailure);
	  });
	}
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


