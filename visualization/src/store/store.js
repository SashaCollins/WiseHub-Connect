import Vue from 'vue';
import VueStore from 'vuex';
import {sidebar} from './sidebar.module'

Vue.use(VueStore);

export default new VueStore.Store({
    modules: {
        sidebar
    }
});