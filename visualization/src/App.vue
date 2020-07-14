<!--   Quelle: [https://github.com/yaminncco/vue-sidebar-menu/blob/master/demo/App.vue]   -->
<template>
  <div
          id="sidebar"
          :class="[{'collapsed' : collapsed}, {'onmobile' : isOnMobile}]"
  >
    <div class="container" id="containerTheme">
      <div>
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
<!--      <hr style="border: 1px solid #e3e3e3;">-->
    </div>
    <router-view />
    <div class="sidebar">
      <sidebar-menu
              :key="selectedTheme"
              :menu="menu"
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
  </div>
</template>

<script>
    // import wisehubIcon from '@/assets/wiesehub-small.svg';
    import Icon from './assets/logo.png';

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
            let loggedIn = false;
            return {
            menu: [
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
                    href:'/repositories',
                    title: 'Repositories',
                    icon: 'fa fa-code fa-fw',
                    hidden: !loggedIn,
                },
                {
                    href: '/courses',
                    title: 'Courses',
                    icon: 'fa fa-chalkboard-teacher fa-fw',
                    hidden: !loggedIn,
                    child: [
                        {
                            href: '/courses/\'vss\'',
                            title: '\'VSS\'',
                            icon: 'fas fa-code-branch fa-fw',
                        },
                    ]
                },
                {
                    href: '/settings',
                    title: 'Settings',
                    icon: 'fas fa-tools fa-fw',
                    hidden: !loggedIn,
                    child: [
                        {
                            href: '/settings/profile',
                            title: 'Profile',
                            icon: 'fas fa-user fa-fw',
                        },
                        {
                            href: '/settings/contact',
                            title: 'Contact',
                            icon: 'fas fa-bullhorn fa-fw',
                        },
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
                    href: '/logout' ,
                    title: 'LogOut',
                    icon: 'fas fa-sign-out-alt fa-fw',
                    hidden: !loggedIn,
                },
                {
                    href: '/login',
                    title: 'LogIn',
                    icon: 'fas fa-sign-out-alt fa-fw',
                    hidden: loggedIn,
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
        // computed () {
        //
        // },
        mounted () {
            this.onResize()
            window.addEventListener('resize', this.onResize)
            window.addEventListener('storage', this.updateTheme)
        },
        methods: {
            updateTheme (event) {
                console.log("event")
                console.log(event)
                if (event.key === 'theme') {
                    console.log(event)
                    this.selectedTheme = '';
                }
            },
            onToggleCollapse (collapsed) {
                console.log(collapsed)
                this.collapsed = collapsed
            },
            onItemClick (event, item) {
                console.log('onItemClick')
                console.log(event)
                console.log(item)
                this.$router.push(item.href)
            },
            onResize () {
                if (window.innerWidth <= 767) {
                    this.isOnMobile = true
                    this.collapsed = true
                } else {
                    this.isOnMobile = false
                    this.collapsed = false
                }
            },
        }
    }
</script>

<style scoped lang="scss">
    @import url('https://fonts.googleapis.com/css?family=Source+Sans+Pro:400,600');
    @import "./scss/sidebar-menu.scss";

    div.position-sticky {
        margin-top: 35px;
        padding: 20px;
        background: #F5FFFA;
        align-content: center;
        position: -webkit-sticky; /* Safari */
        position: sticky;
        top: 0;
    }
    .container-fluid {
        width: 75%;
        height: fit-content;
        margin-top: 35px;
        padding: 20px;
        border-radius: 10px;
        background: #F5FFFA;
        align-content: center;
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


</style>

