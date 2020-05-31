import AuthService from '../services/auth.service';
import jwtDecode from "jwt-decode";

const user = sessionStorage.getItem('shared');
const initialState = user
    ? { status: { loggedIn: true }, user }
    : { status: { loggedIn: false }, user: null };

export const auth = {
    namespaced: true,
    state: initialState,
    actions: {
        login({ commit }, user) {
            return AuthService.login(user).then(
                response => {
                    let decoded = jwtDecode(response.data.access)
                    user.id = decoded.identity.user_ref;
                    user.username = decoded.identity.user_name;
                    user.password = '';
                    commit('loginSuccess', user);
                    return Promise.resolve(user);
                },
                error => {
                    commit('loginFailure');
                    return Promise.reject(error);
                }
            );
        },
        refresh({ commit }) {
            return AuthService.refresh().then(
                response => {
                    if (response.data) {
                        console.log(response);
                        commit('refreshSuccess');
                    }
                    return Promise.resolve(response);
                },
                error => {
                    commit('refreshFailure');
                    return Promise.reject(error);
                }
            );
        },
        logout({ commit }) {
            AuthService.logout();
            commit('logout');
        },
        register({ commit }, user) {
            return AuthService.register(user).then(
                response => {
                    user.password = '';
                    commit('registerSuccess');
                    return Promise.resolve(response.data);
                },
                error => {
                    commit('registerFailure');
                    return Promise.reject(error);
                }
            );
        },
        update_profile({ commit }, user) {
            return UserService.update(user).then(response => {
                commit("updateSuccess", user)
                return Promise.resolve(response);
            }, error => {
                commit("updateFailure")
                return Promise.reject(error);
            });
        },
        update_password({ commit }, user) {
            return UserService.update(user).then(response => {
                user.password = '';
                commit("updateSuccess", user)
                return Promise.resolve(response);
            }, error => {
                commit("updateFailure")
                return Promise.reject(error);
            });
        },
        delete({ commit }, user) {
            return UserService.delete(user).then(response => {
                commit("deleteSuccess")
                user.password = '';
                sessionStorage.clear();
                localStorage.clear();
                return Promise.resolve(response);
            }, error => {
                commit("deleteFailure")
                return Promise.reject(error);
            });
        }
    },
    mutations: {
        loginSuccess(state, user) {
            state.status.loggedIn = true;
            state.user = user;
        },
        loginFailure(state) {
            state.status.loggedIn = false;
            state.user = null;
        },
        logout(state) {
            state.status.loggedIn = false;
            state.user = null;
        },
        registerSuccess(state) {
            state.status.loggedIn = false;
        },
        registerFailure(state) {
            state.status.loggedIn = false;
        },
        updateSuccess(state, user) {
            state.status.loggedIn = true;
            state.user = user;
        },
        updateFailure(state) {
            state.status.loggedIn = true;
        },
        deleteSuccess(state) {
            state.status.loggedIn = false;
            state.user = null;
        },
        deleteFailure(state) {
            state.status.loggedIn = true;
        }
    }
};