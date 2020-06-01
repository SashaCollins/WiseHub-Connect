import SidebarService from '../services/sidebar.service';

let theme = localStorage.getItem('theme');
const initialState = theme
    ? { status: { exists: true }, theme }
    : { status: { exists: false }, theme: '' };

export const sidebar = {
    namespaced: true,
    initialState,
    actions: {
        theme({ commit }, selected) {
            SidebarService.update(selected)
            // set store
            commit('update', selected);
        },
    },
    mutations: {
        update(state, selected) {
            state.exists = true;
            state.theme = selected;
            console.log(state)
        },
    }
};