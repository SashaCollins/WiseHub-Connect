export const sidebar = {
    namespaced: true,
    state: {
        selectedTheme: 'wisehub-theme',
    },
    actions: {
        theme({ commit }, selectedTheme) {
            // set storage
            localStorage.setItem('selectedTheme', selectedTheme);
            // set store
            commit('updateTheme', selectedTheme);
        },
    },
    mutations: {
        updateTheme(state, selectedTheme) {
            state.selectedTheme = selectedTheme;
        },
    }
};