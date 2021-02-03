<template>
  <div class="container-fluid">
    <div class="header">
      <h2>WiseHub-Courses</h2>
    </div>
    <div v-if="!message">
      <div v-if="!loading">
        <VersionManagerCards v-bind:vmcards="getVersionManagerCards('Github')"></VersionManagerCards>
      </div>
      <div v-else>
        <div class="spinner-border text-info" role="status">
          <span class="sr-only">Loading...</span>
        </div>
      </div>
    </div>
    <div v-else>
      {{ this.message }}
    </div>

    <!--      load this part only in 'onItemClick'-->
    <div class="row" v-if="clicked">
      <div
          id="accordion"
          v-for="(item, index) in plugins"
          :key="index"
          class="col-12">

        <div class="card">
          <div class="card-header" :id="'h' + index">
            <h5 class="mb-0">
              <button class="btn btn-info" data-toggle="collapse" :data-target="'#c' + index" aria-expanded="true" :aria-controls="'c' + index">
                {{ item.PluginName }}
              </button>
            </h5>
          </div>

          <div :id="'c' + index" class="collapse show" :aria-labelledby="'h' + index" data-parent="#accordion">
            <div class="card-body">
              {{ item.PluginContent }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
    import VersionManagerCards from "@/components/helper/VersionManagerCards";

    export default {
      name: "VersionManager",
      components: {
        VersionManagerCards,
      },
      data() {
        return {
          clicked: false,
          message: '',
          loading: false,
        }
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
        getVersionManagerCards (name) {
          for (let i = 0; i < this.getUser.plugins.length; i++) {
            let plugin = this.getUser.plugins[i];
            if (name === plugin.pluginName) {
              return plugin;
            }
          }
        }
      },
      mounted() {
        this.loading = true;
        if (this.getUser.plugins.length === 0) {
          this.$store.dispatch('user/fetchData', {
            option: "general",
            user: this.getUser,
          }).then(
              (onSuccess) => {
                if (onSuccess.data.success) {
                  if (onSuccess.data.pluginData) {
                    this.plugins = this.getUser.plugins;
                  }
                }
              },
              (onError) => {
                this.message = onError.toString() || onError.message;
              }
          );
        }
        this.loading = false;
      },
    }
</script>

<style scoped lang="scss">
  .row {
    width: 100%;
    padding-left: 10px;
    //border-radius: 10px;
  }
  .row:after {
    content: "";
    display: table;
    clear: both;
  }

  .col {
    //border-radius: 10px;
    //margin-top: 20px;
  }
  /* Add a card effect for articles */
  .card {
    float: top;
    background-color: white;
    //padding: 20px;
    //margin-top: 20px;
  }

  /* If you want the content centered horizontally and vertically */
  .centered {
    position: relative;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    text-align: center;
  }
</style>