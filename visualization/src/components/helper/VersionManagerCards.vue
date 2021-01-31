<template>
  <div class="container-fluid">
    <div class="row">
      <div v-for="(orga, index) in vmcards.pluginData" :key="index">
        <h2 class="text-center" style="background-color: #464646; color: white; border-radius: 3px; padding: 10px">
          Organisation: {{orga.organization.orgaName}}
        </h2>
        <div class="card-deck">
          <div v-for="(team, index) in orga.organization.teams" :key="index">
            <div class="card">
              <h3 class="text-center" style="background-color: #464646; color: white; border-radius: 3px; padding: 10px">
                Team: {{ team.teamName }}
              </h3>
              <div v-for="(repo, index) in team.repositories" :key="index">
                <h4 class="text-center" style="background-color: #464646; color: white; border-radius: 3px; padding: 10px">
                  Repository: {{repo.repoName}}
                </h4>
                <div class="card bg-white text-center">
                  <Github v-bind:github="team"></Github>
                  <DroneCI v-bind:drone="getRepo('Drone CI', repo.repoName)"></DroneCI>
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
import Heroku from "@/components/plugins/Heroku";

export default {
  name: "VersionManagerCards",
  props: ["vmcards"],
  components: {
    Github,
    DroneCI,
    Heroku,
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
    getRepo (pluginName, repoName) {
      for (let i = 0; i < this.getUser.plugins.length; i++) {
        let plugin = this.getUser.plugins[i];
        if (pluginName === plugin.pluginName) {
          for (let j = 0; j < plugin.pluginData.length; j++) {
            let data = plugin.pluginData[j];
            console.log(data);
            console.error(repoName === data.repo.repoName)
            if (repoName.trim() === data.repo.repoName.trim()) {
              console.error(repoName);
              console.error(data);
              return data;
            }
          }
        }
      }
    }
  },
  // data() {
  //   return {
  //     github: [{
  //       'Name': '',
  //       'Teams': [{
  //         'Name': '',
  //         'Members': [{
  //           'Name': '',
  //         }],
  //         'Repositories': [{
  //           'Name': '',
  //           'URL': '',
  //           'Issues': [{
  //             'Name': '',
  //             'Titel': '',
  //           }],
  //         }],
  //       }],
  //     }],
  //   }
  // },
}
</script>

<style scoped>

</style>