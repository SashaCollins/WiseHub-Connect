import Vue from 'vue';
import Router from 'vue-router';

import Homepage from '../components/hub/Homepage';
import FAQ from '../components/info/FAQ.vue';
import Profile from "../components/profile/Profile"
import Contact from "../components/info/Contact"
import Impressum from "../components/info/Impressum";
import LogIn from '../components/auth/LogIn.vue';
import SignUp from '../components/auth/SignUp.vue';
import Forgot from '../components/auth/Forgot.vue';
import TemplateView from "../components/hub/TemplateView";
import GeneralView from "../components/hub/GeneralView";
import { secure } from '@/services/encryption.service';

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
           name: 'templateview',
           path: '/view/templateview',
           component: TemplateView,
       },
       {
           name: 'generalview',
           path: '/view/generalview',
           component: GeneralView,
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
           path: '/impressum',
           component: Impressum,
       },
       {
           name: 'logout',
           path: '/logout',
       },
       {
           name: 'login',
           path: '/login',
           component: LogIn,
       },
       {
           name: 'signup',
           path: '/signup',
           component: SignUp,
       },
       {
           name: 'forgot',
           path: '/forgot',
           component: Forgot,
       },
      
      // otherwise redirect to home
      { path: '*', redirect: '/' },
   ],
   linkActiveClass: 'active',
});

router.beforeEach((to,from,next) => {
   let publicPages = ['/','/login','/faq','/impressum','/signup','/forgot', '/validate'];
   let authRequired = !publicPages.includes(to.path);
   const loggedIn = secure.get('token');

   if (authRequired && !loggedIn) {
       next('/login');
   }
   next();
})
