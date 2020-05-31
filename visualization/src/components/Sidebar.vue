<!--   Quelle: [https://github.com/yaminncco/vue-sidebar-menu/blob/master/demo/App.vue]   -->
<template>
    <div
            id="sidebar"
            :class="[{'collapsed' : collapsed}, {'onmobile' : isOnMobile}]"
    >
        <router-view />
        <div class="sidebar">
            <sidebar-menu
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
        name: "Sidebar",
        data() {
            return {
                menu: [
                    {
                        header: true,
                        title: 'WiseHub Dashboard',
                        hiddenOnCollapse: true
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
                    },
                    {
                        href: '/courses',
                        title: 'Courses',
                        icon: 'fa fa-chalkboard-teacher fa-fw',
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
                            {
                                href: '/settings/impressum',
                                title: 'Impressum',
                                icon: 'fas fa-copyright fa-fw',
                            },
                        ]
                    },
                    {
                        component: separator
                    },
                    {
                        href: { path: '/logout' },
                        title: 'LogOut',
                        icon: 'fas fa-sign-out-alt fa-fw',
                    },
                ],
                collapsed: false,
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
                isOnMobile: false
            }
        },
        mounted () {
            this.onResize()
            window.addEventListener('resize', this.onResize)
        },
        methods: {
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
            }
        }
    }
</script>

<style scoped lang="scss">
    @import url('https://fonts.googleapis.com/css?family=Source+Sans+Pro:400,600');
    @import "../scss/sidebar-menu.scss";

    #sidebar {
        padding-left: 350px;
        transition: 0.3s ease;
    }
    #sidebar.collapsed {
        padding-left: 350px;
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
    .container {
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