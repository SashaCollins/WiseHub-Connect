import AuthService from '../services/auth.service';

const loggedIn = sessionStorage.getItem('loggedIn');
const initialState = loggedIn
    ? { status: {loggedIn: true}}
    : { status: {loggedIn: false}};

export const auth = {
    namespaced: true,
    state: initialState,
    actions: {
        login({ commit }, user) {
            return AuthService.login(user).then(
                (onSuccess) => {
                    if (onSuccess.data.success){
                        commit('loginSuccess');
                    }
                    return Promise.resolve(onSuccess);
                },
                (onFailure) => {
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
                    return Promise.resolve(onSuccess.data);
                },
                onFailure => {
                    return Promise.reject(onFailure);
                }
            );
        },
    },
    mutations: {
        loginSuccess(state) {
            state.status.loggedIn = true;
        },
        logout(state) {
            state.status.loggedIn = false;
        },
    }
};
