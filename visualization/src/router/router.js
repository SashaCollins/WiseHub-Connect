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
   ]
});

router.beforeEach((to,from,next) => {
    let publicPages = ['/','/login','/faq','/impressum','/signup','/forgot'];
    let authRequired = !publicPages.includes(to.path);
    let loggedIn = sessionStorage.getItem('loggedIn');
    if (authRequired && !loggedIn) {
        return next('/login');
    }
    next();
})
