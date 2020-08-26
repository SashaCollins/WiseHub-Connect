import AuthService from '../services/auth.service';

const loggedIn = sessionStorage.getItem('loggedIn');
const initialState = loggedIn
    ? { status: {loggedIn: loggedIn}}
    : { status: {loggedIn: loggedIn}};

export const auth = {
    namespaced: true,
    state: initialState,
    actions: {
        login({ commit }, user) {
            return AuthService.login(user).then(
                (onSuccess) => {
                    if (onSuccess.data.Success){
                        commit('loginSuccess');
                    }
                    return Promise.resolve(onSuccess);
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
        loginSuccess(state) {
            state.status.loggedIn = true;
        },
        loginFailure(state) {
            state.status.loggedIn = false;
        },
        logout(state) {
            state.status.loggedIn = false;
        },
        registerSuccess(state) {
            state.status.loggedIn = false;
        },
        registerFailure(state) {
            state.status.loggedIn = false;
        },
    }
};