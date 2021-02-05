<!--   Quelle: [https://github.com/yaminncco/vue-sidebar-menu/blob/master/demo/App.vue]   -->
<template>
    <div
            id="sidebar"
            :class="[{'collapsed' : collapsed}, {'onmobile' : isOnMobile}]"
    >
      <router-view />
      <div class="sidebar">
        <sidebar-menu
                :key="selectedTheme"
                :menu="loggedIn ? menuLoggedIn : menuLoggedOut"
                :collapsed="collapsed"
                :theme="selectedTheme"
                :show-one-child="true"
                :show-child="false"
                @toggle-collapse="onToggleCollapse"
                @item-click="onItemClick"
        />
        <div
                v-if="isOnMobile && !collapsed"
                class="sidebar-overlay"
                @click="collapsed = true"
        />
      </div>

      <div class="theme">
        <!--        {{ loggedIn }}:-->
        <select v-model="selectedTheme">
          <option
              v-for="(theme, index) in themes"
              :key="index"
              :value="theme.input"
          >
            {{ theme.name }}
          </option>
        </select>
      </div>
    </div>
</template>

<script>
    import Icon from './assets/wisehubIcon.png';

    const separator = {
        render (h) {
              return h('hr', {
                    style: {
                        borderColor: 'rgba(0, 0, 0, 0.1)',
                        margin: '20px'
                    }
              })
        }
    }

    export default {
        name: "App",
        data() {
            let loggedIn = this.loggedIn;
            return {
            menuLoggedOut: [
                {
                    header: true,
                    title: 'WiseHub Dashboard',
                    hiddenOnCollapse: true
                },
                {
                    href: '/',
                    title: 'Homepage',
                    icon: {
                        element: 'img',
                        attributes: {
                            src: Icon,
                        }
                    },
                },
                {
                    href: '/faq',
                    title:  'FAQ',
                    icon: 'far fa-comments fa-fw',
                },
                {
                    href: '/impressum',
                    title: 'Impressum',
                    icon: 'fas fa-copyright fa-fw',
                },
                {
                    component: separator
                },
                {
                    href: '/login',
                    title: 'LogIn',
                    icon: 'fas fa-sign-out-alt fa-fw',
                    hidden: loggedIn,
                },
            ],
            menuLoggedIn: [
              {
                header: true,
                title: 'WiseHub Dashboard',
                hiddenOnCollapse: true
              },
              {
                href: '/',
                title: 'Homepage',
                icon: {
                  element: 'img',
                  attributes: {
                    src: Icon,
                  }
                },
              },
              {
                href: '/faq',
                title:  'FAQ',
                icon: 'far fa-comments fa-fw',
              },
              {
                href:'/view/templateview',
                title: 'TemplateView',
                icon: 'fa fa-code fa-fw',
              },
              {
                href: '/view/generalview',
                title: 'GeneralView',
                icon: 'fa fa-chalkboard-teacher fa-fw',
              },
              {
                title: 'Settings',
                icon: 'fas fa-tools fa-fw',
                child: [
                  {
                    href: '/settings/profile',
                    title: 'Profile',
                    icon: 'fas fa-user fa-fw',
                  },
                    // Activate contact  form to enable sending contact messages
                  // {
                  //   href: '/settings/contact',
                  //   title: 'Contact',
                  //   icon: 'fas fa-bullhorn fa-fw',
                  // },
                ]
              },
              {
                href: '/impressum',
                title: 'Impressum',
                icon: 'fas fa-copyright fa-fw',
              },
              {
                component: separator
              },
              {
                href: '/logout',
                title: 'LogOut',
                icon: 'fas fa-sign-out-alt fa-fw',
              },
            ],
            themes: [
                {
                    name: 'WiseHub theme',
                    input: 'wisehub-theme'
                },
                {
                    name: 'Black theme',
                    input: 'black-theme'
                }
            ],
            selectedTheme: 'wisehub-theme',
            collapsed: false,
            isOnMobile: false,
          }
        },
        computed: {
          //switch menu between user is loggedIn and user is loggedOut
          loggedIn() {
            return this.$store.state.auth.status.loggedIn;
          },
        },
        mounted () {
            this.onResize();
            window.addEventListener('resize', this.onResize);
        },
        methods: {
            onToggleCollapse (collapsed) {
                this.collapsed = collapsed;
            },
            onItemClick (event, item) {
                if (item.href === '/logout') {
                  this.$store.dispatch("auth/logout").then(
                      () => {
                        this.$router.push("/");
                      },
                      (onFailure) => {
                        this.message = onFailure.response.data;
                        this.submitted = false;
                      })
                }
                if (item.href !== ''){
                  this.$router.push(item.href);
                }
            },
            onResize () {
                if (window.innerWidth <= 767) {
                    this.isOnMobile = true;
                    this.collapsed = true;
                } else {
                    this.isOnMobile = false;
                    this.collapsed = false;
                }
            },
        }
    }
</script>

<style lang="scss">
    @import url('https://fonts.googleapis.com/css?family=Source+Sans+Pro:400,600');
    @import "./scss/sidebar-menu.scss";

    * {
      box-sizing: border-box;
    }
    .bg-img {
      /* The image used */
      background-image: url("./assets/wisehubLogoV.png");

      /* Control the height of the image */
      min-height: 100vh;

      /* Center and scale the image nicely */
      background-position: center;
      background-repeat: no-repeat;
      background-size: cover;
      position: relative;
    }

    div.header {
        margin-top: 35px;
        background: #F5FFFA;
        //align-content: center;
        position: -webkit-sticky; /* Safari */
        position: sticky;
        top: 0;
        width: 100%;
        padding-top: 20px;
        padding-bottom: 20px;
        text-align: center;
        z-index: 1;
    }
    h2 {
      color: #008B8B;
      width: 100%;
      margin-top: 15px;
      border-radius: 10px;
      text-align: center;
      background: #F5FFFA;
    }

    @media screen and (max-width: 767px) {
        .container-fluid {
            width: 100%;
        }
    }

    #sidebar {
        padding-left: 350px;
        transition: 0.3s ease;
    }
    #sidebar.collapsed {
        padding-left: 75px;
    }
    #sidebar.onmobile {
        padding-left: 50px;
    }
    .sidebar-overlay {
        position: fixed;
        width: 100%;
        height: 100%;
        top: 0;
        left: 0;
        background-color: #000;
        opacity: 0.5;
        z-index: 900;
    }
    .sidebar {
        padding: 50px;
    }
    #containerTheme {
        max-width: 1900px;
    }
    pre {
        font-family: Consolas, monospace;
        color: #000;
        background: #fff;
        border-radius: 2px;
        padding: 15px;
        line-height: 1.5;
        overflow: auto;
    }
    .theme {
      //position: -webkit-sticky; /* Safari */
      position: fixed;
      bottom: 0;
    }


</style>

