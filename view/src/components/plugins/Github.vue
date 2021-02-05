<template>
  <div class="card">
    {{ this.github}}
    <div class="card-header">
      <h5
          class="text-center"
          style="background-color: #008B8B; color: white; border-radius: 3px; padding: 15px">
        Github
      </h5>
    </div>
    <div class="card-body text-left">
      <table class="card-table table table-bordered">
        <thead>
        <tr>
          <th>Members</th>
          <th>Your Issues</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(member, index) in github.members" :key="index">
          <td>{{ member.memberName }}</td>
        </tr>
        <tr v-for="(issue, index) in github.issues" :key="index">
          <td>
            <p>ID: #{{issue.issueNumber}}</p>
            <p>Title: {{issue.issueTitle}}</p>
          </td>
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
    },
  },
}
</script>

<style scoped>

</style>