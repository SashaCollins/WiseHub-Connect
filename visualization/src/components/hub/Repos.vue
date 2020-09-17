<template>
    <div class="container">
      <div class="header">
        <h2>WiseHub-Repositories</h2>
      </div>

      <div class="row">
        <div
            v-for="(item, index) in repos"
            :key="index"
            class="col-lg-6 col-md-12 col-sm-12">
          <div
              class="card"
              v-on:click="onItemClick">
            <h3
                class="text-center"
                style="background-color: #464646; color: white; border-radius: 3px; padding: 15px">
              {{ item.RepoName }}
            </h3>
<!--            <div class="card-body" style="padding: 0; margin: 0;">-->
              <div class="row" style="padding-left: 30px; padding-bottom: 10px;">
              <div class="col" style="background-color: #283230; color: white; border-radius: 3px;">
                <div class="centered">
                  <label
                      :for="item.Contributors"
                      class="col-md col-form-label">
                    Contributors:
                  </label>
                  <ul class="list-group">
                    <li v-for="item in item.Contributors" :key="item">
                      {{ item }}
                    </li>
                  </ul>
                </div>
              </div>
              <div class="col offset-1" style="background-color: #283230; color: white; border-radius: 3px;">
                <div class="centered">
                  <label
                      :for="item.Description"
                      class="col-md col-form-label">
                    Description:
                  </label>
                  {{ item.Description }}
                </div>
              </div>
              </div>
          </div>

        </div>
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
<!--        <div-->
<!--            v-for="(dings, index) in plugins"-->
<!--            :key="index"-->
<!--            class="col-12">-->
<!--          <div-->
<!--              class="card">-->
<!--            <h3-->
<!--                class="text-center"-->
<!--                style="background-color: #464646; color: white; border-radius: 3px; padding: 15px">-->
<!--              {{ dings.PluginName }}-->
<!--            </h3>-->
<!--            <div class="centered">-->
<!--              <label-->
<!--                  :for="dings.PluginContent"-->
<!--                  class="col-md col-form-label">-->
<!--              </label>-->
<!--              {{ dings.PluginContent }}-->
<!--            </div>-->
<!--          </div>-->
<!--        </div>-->
      </div>
</template>

<script>
    export default {
      name: "Repos",
      data() {
        return {
          clicked: false,
          //dummy only exists if connection fails
          repos: [{
            'RepoName': 'Dummy',
            'Contributors': ['Hans Wurst', 'Axel Schweiß'],
            'Description': 'This Repo is not real fetch data from GR',
          }],
          plugins: [{
            'PluginName': 'DummyPlugin1',
            'PluginContent': 'DummyContent'
          },{
            'PluginName': 'DummyPlugin2',
            'PluginContent': 'DummyContent'
          }]
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
        onItemClick: function (event, item) {
          this.clicked = true;
          console.log('onItemClick')
          console.log(event)
          console.log(item)
        },
        fetchAllRepos: function() {
          this.$store.dispatch('user/fetchRepos', this.getUser).then(
              (onSuccess) => {

          }, (onError) => {

          })
        },
        fetchRepo: function(item) {

        }
      },
      mounted() {
        console.log(this.getUser);
        this.$store.dispatch('user/fetchRepos', this.getUser).then(
            (onSuccess) => {
              if (onSuccess.data.success) {
                // this.plugins = onSuccess.data.plugins;
              }
            },
            (onError) => {
              // this.message = onError.toString() || onError.message;
            }
        )
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