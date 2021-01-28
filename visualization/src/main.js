import Vue from 'vue';
import App from './App.vue';
import VueResource from 'vue-resource';
import * as VeeValidate from 'vee-validate';

import 'bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
// import BootstrapVue from 'bootstrap-vue/dist/bootstrap-vue.esm';
import {BootstrapVue, IconsPlugin, vue} from 'bootstrap-vue'
import { LayoutPlugin } from 'bootstrap-vue'
Vue.use(LayoutPlugin)
import { BContainer, BRow, BCol } from 'bootstrap-vue'
Vue.component('b-container', BContainer)
Vue.component('b-row', BRow)
Vue.component('b-col', BCol)

import 'vue-sidebar-menu/dist/vue-sidebar-menu.css'
import VueSidebarMenu from 'vue-sidebar-menu'
import SidebarMenu from './scss/sidebar-menu.scss'

import { router } from './router/router';
import store from './store/index';

import { library, dom } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { fas } from '@fortawesome/free-solid-svg-icons';
import '@fortawesome/fontawesome-free/css/all.css';
import '@fortawesome/fontawesome-free/js/all.js';

dom.watch();
library.add(fas);

Vue.config.productionTip = false;

Vue.use(VueSidebarMenu);
Vue.use(SidebarMenu);
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
Vue.use(VeeValidate);
Vue.use(VueResource);
Vue.component('font-awesome-icon', FontAwesomeIcon);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');