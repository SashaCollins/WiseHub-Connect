import Vue from 'vue';
import Router from 'vue-router';
import Homepage from '../components/Homepage';
import FAQ from '../components/FAQ.vue';
import Profile from "../components/Profile"
import Contact from "../components/Contact"
import Impressum from "../components/Impressum";
import LogIn from '../components/LogIn.vue';
Vue.use(Router);

export const router = new Router({
   mode: 'history',
   routes: [
       {
           name: 'homepage',
           path: '/',
           component: Homepage,
       },
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
           path: '/settings/contact',
           component: Contact,
       },
       {
           name: 'impressum',
           path: '/settings/impressum',
           component: Impressum,
       },
       {
           name: 'logout',
           path: '/logout',
           // function: logout(),
       },
       {
           name: 'login',
           path: '/login',
           component: LogIn,
       },
   ]
});