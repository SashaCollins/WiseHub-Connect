<template>
  <div class="card">
    <div class="card-header">
      <h5
          class="text-center"
          style="background-color: #008B8B; color: white; border-radius: 3px; padding: 15px">
        Github
      </h5>
    </div>
    <div class="card-body text-left" style="overflow-y: auto; min-width: fit-content;">
      <div>
        Members:
        <div class="col-12" v-for="(member, index) in githubTeam.members" :key="index">
          {{ member.memberName }}
        </div>
      </div>
      <div>
        My Issues:
        <div class="col-12" v-for="(issue, index) in githubRepo.issues" :key="index">
          ID: #{{issue.issueNumber}}
          Title: {{issue.issueTitle}}
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "Github",
  props: ["githubTeam", "githubRepo"],
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