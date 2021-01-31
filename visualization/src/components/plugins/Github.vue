<template>
  <div class="card">
    <div class="card-header">
      <h5
          class="text-center"
          style="background-color: #464646; color: white; border-radius: 3px; padding: 15px">
        Github
      </h5>
    </div>
    <div class="card-body text-left">
      <table class="card-table table table-bordered">
        <thead>
        <tr>
          <th>Members</th>
          <th>Issues</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(member, index) in github.members" :key="index">
          <td>{{ member.memberName }}</td>
        </tr>
        <tr v-for="(issue, index) in github.issues" :key="index">
          <td>{{ issue }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>

export default {
  name: "Github",
  props: ["github"],
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    },
    getUser() {
      return this.$store.state.user.user;
    }
  },
  methods: {
    onItemClick (event, item) {
      this.clicked = true;
      // console.log('onItemClick');
      // this.$store.dispatch('user/fetchTeams', {
      //   user: this.getUser,
      //   organization: course
      // }).then((onSuccess) => {
      //   this.plugins = onSuccess.data.plugins;
      // }, (onError) => {
      //   this.message = onError.toString() || onError.message;
      // });
    },
  },
}
</script>

<style scoped>

</style>