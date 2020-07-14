import Vue from 'vue';
import Vuex from 'vuex';
import { sidebar } from './sidebar.module'
import { auth } from './auth.module.js'

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        sidebar,
        auth,

    }
});