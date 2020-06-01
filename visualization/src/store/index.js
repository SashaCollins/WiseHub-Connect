import Vue from 'vue';
import Vuex from 'vuex';
import { sidebar } from './sidebar.module'

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        sidebar
    }
});