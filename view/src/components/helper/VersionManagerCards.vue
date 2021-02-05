<template>
  <div class="container-fluid">
    <div class="row">
      <div v-for="(orga, index) in vmCards.pluginData" :key="index">
        <h2 class="text-center" style="background-color: #008b8b; color: white; border-radius: 3px; padding: 10px">
          Organisation: {{orga.organization.orgaName}}
        </h2>
        <div class="card-columns">
          <div v-for="(team, index) in orga.organization.teams" :key="index">
            <div class="card">
              <h5 class="text-center" style="background-color: white; color: black; border-radius: 3px; padding: 10px">
                Team: {{ team.teamName }}
              </h5>
              <div v-for="(repo, index) in team.repositories" :key="index">
                <h5 class="text-center" style="background-color: white; color: black; border-radius: 3px; padding: 10px">
                  Repository: {{repo.repoName}}
                  <br>
                  URL: <a :href="repo.repoUrl" >{{repo.repoName}}</a>
                </h5>
                <div class="card bg-white text-center">
                  <Github v-bind:githubTeam="team" v-bind:githubRepo="getGithubRepo(team, repo.repoName)"></Github>
                  <DroneCI v-bind:drone="getDroneCIRepo('Drone CI', repo.repoName)"></DroneCI>
<!--                  <Template v-bind:template="getRepo('TemplateName', repo.repoName)"></Template>-->
                </div>
              </div>
            </div>
          </div>
        </div>
        <br>
      </div>
    </div>
  </div>
</template>

<script>
import Github from "@/components/plugins/Github";
import DroneCI from "@/components/plugins/DroneCI";
// import Template from "@/components/plugins/Template";

export default {
  name: "VersionManagerCards",
  props: ["vmCards"],
  components: {
    Github,
    DroneCI,
    // Template,
  },
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
    //combines repo information from all plugins based on the repoName
    getDroneCIRepo (pluginName, repoName) {
      for (let i = 0; i < this.getUser.plugins.length; i++) {
        let plugin = this.getUser.plugins[i];
        if (pluginName === plugin.pluginName) {
          for (let j = 0; j < plugin.pluginData.length; j++) {
            let data = plugin.pluginData[j];
            if (repoName.trim() === data.repositories.repoName.trim()) {
              return data;
            }
          }
        }
      }
    },
    getGithubRepo (githubTeam, repoName) {
      for (let i = 0; i < githubTeam.repositories.length; i++) {
        let repo = githubTeam.repositories[i];
        if (repoName === repo.repoName) {
          return repo;
        }
      }
    }
  },
}
</script>

<style scoped>

</style>