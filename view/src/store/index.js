import Vue from 'vue';
import Vuex from 'vuex';
import { auth } from './auth.module.js'
import { user } from './user.module.js'

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        auth,
        user
    }
});