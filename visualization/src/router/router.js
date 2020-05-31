import Vue from 'vue';
import Router from 'vue-router';
import FAQ from '../components/FAQ.vue';
import Profile from "../components/Profile"
import Contact from "../components/Contact"
import Impressum from "../components/Impressum";
import LogOut from '../components/LogOut.vue';
Vue.use(Router);

export const router = new Router({
   mode: 'history',
   routes: [
       {
           name: 'faq',
           path: '/faq',
           component: FAQ,
       },
       {
           name: 'settings',
           path: '/settings',
       },
       {
           name: 'profile',
           path: '/settings/profile',
           component: Profile,
       },
       {
           name: 'contact',
           path: '/contact',
           component: Contact,
       },
       {
           name: 'impressum',
           path: '/impressum',
           component: Impressum,
       },
       {
           name: 'logout',
           path: '/logout',
           component: LogOut,
       },
   ]
});