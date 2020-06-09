<template>
    <div class="container">
        <div class="position-sticky">
            <h2>WiseHub-Profile</h2>
        </div>
        <b-container fluid="true" class="container-fluid">
<!--            <h4>Hi {{account.user.firstName}}!</h4>-->
<!--            <p>You're logged in with Vue + Vuex & JWT!!</p>-->
<!--            <h6>Users from secure api end point:</h6>-->
<!--            <em v-if="users.loading">Loading users...</em>-->
<!--            <span v-if="users.error" class="text-danger">ERROR: {{users.error}}</span>-->
<!--            <ul v-if="users.items">-->
<!--                <li v-for="user in users.items" :key="user.id">-->
<!--                    {{user.firstName + ' ' + user.lastName}}-->
<!--                    <span v-if="user.deleting"><em> - Deleting...</em></span>-->
<!--                    <span v-else-if="user.deleteError" class="text-danger"> - ERROR: {{user.deleteError}}</span>-->
<!--                    <span v-else> - <a @click="deleteUser(user.id)" class="text-danger">Delete</a></span>-->
<!--                </li>-->
<!--            </ul>-->
        </b-container>

    </div>
</template>

<script>
    import EventBus from '../bus/event.bus';
    import { mapState, mapActions } from 'vuex'

    export default {
        name: "Profile",
        computed: {
            ...mapState({
                account: state => state.account,
                users: state => state.users.all
            })
        },
        created () {
            this.getAllUsers();
        },
        methods: {
            ...mapActions('users', {
                getAllUsers: 'getAll',
                deleteUser: 'delete'
            })
        }
    }
</script>

<style lang="scss">
    h2 {
        color: #008B8B;
        width: 100%;
        margin-top: 15px;
        border-radius: 10px;
        text-align: center;
        background: #F5FFFA;
    }

    @media screen and (max-width: 800px) {

    }
</style>