import AuthService from '../services/auth.service';
import jwtDecode from "jwt-decode";

const loggedIn = sessionStorage.getItem('loggedIn');
const user = sessionStorage.getItem('user');
const initialState = loggedIn
    ? { status: {loggedIn: loggedIn, user: user}}
    : { status: {loggedIn: loggedIn, user: null }};

export const auth = {
    namespaced: true,
    // { status: loggedIn, user: user }
    state: initialState,
    actions: {
        login({ commit }, user) {
            return AuthService.login(user).then(
                (onSuccess) => {
                    if (onSuccess.data.Success){
                        commit('loginSuccess', user);
                    }
                    return Promise.resolve(user);
                },
                (onFailure) => {
                    commit('loginFailure');
                    return Promise.reject(onFailure);
                }
            );
        },
        logout({ commit }) {
            AuthService.logout();
            commit('logout');
        },
        register({ commit }, user) {
            return AuthService.register(user).then(
                onSuccess => {
                    user.password = '';
                    commit('registerSuccess');
                    return Promise.resolve(onSuccess.data);
                },
                onFailure => {
                    commit('registerFailure');
                    return Promise.reject(onFailure);
                }
            );
        },
    },
    mutations: {
        loginSuccess(state, user) {
            state.status.loggedIn = true;
            state.status.user = user;
        },
        loginFailure(state) {
            state.status.loggedIn = false;
            state.status.user = null;
        },
        logout(state) {
            state.status.loggedIn = false;
            state.status.user = null;
        },
        registerSuccess(state) {
            state.status.loggedIn = false;
        },
        registerFailure(state) {
            state.status.loggedIn = false;
        },
    }
};