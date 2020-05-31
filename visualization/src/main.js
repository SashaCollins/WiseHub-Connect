import Vue from 'vue';
import App from './App.vue';
import VueResource from 'vue-resource';
import BootstrapVue from 'bootstrap-vue/dist/bootstrap-vue.esm';
import 'bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import 'vue-sidebar-menu/dist/vue-sidebar-menu.css'
import VueSidebarMenu from 'vue-sidebar-menu'
import SidebarMenu from './scss/sidebar-menu.scss'
import VeeValidate from 'vee-validate';
import { router } from './router/router';
import store from './store/store';
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import {
  faHome,
  faUser,
  faUserPlus,
  faSignInAlt,
  faSignOutAlt,
  faCoffee, faSpinner, faEdit, faCircle, faCheck,
  faPlus, faEquals, faArrowRight, faPencilAlt, faComment,
} from '@fortawesome/free-solid-svg-icons';
import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'

library.add(faHome, faUser, faUserPlus, faSignInAlt, faSignOutAlt, faCoffee, faSpinner, faEdit, faCircle, faCheck,
    faPlus, faEquals, faArrowRight, faPencilAlt, faComment,);

Vue.config.productionTip = false;

Vue.use(VueSidebarMenu)
// Vue.use(SidebarMenu)
Vue.use(BootstrapVue);
Vue.use(VeeValidate);
Vue.use(VueResource);
Vue.component('font-awesome-icon', FontAwesomeIcon);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');