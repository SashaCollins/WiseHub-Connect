<template>
  <div class="container">
    <div class="header">
      <h2>WiseHub-Courses</h2>
    </div>

    <div class="row">
      <div
          v-for="(item, index) in courses"
          :key="index"
          class="col-lg-6 col-md-12 col-sm-12">
        <div class="card">
          <h3
              class="text-center"
              style="background-color: #464646; color: white; border-radius: 3px; padding: 15px">
            {{ item.Name }}
          </h3>

            <div
                v-for="(item, index) in courses['Teams']"
                :key="index"
                class="col-lg-6 col-md-12 col-sm-12">
              <div class="card">
                <h3 class="text-center">{{ item.Name }}</h3>
                <div class="card-body">

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
      name: "Courses",
      data() {
        let teams;
        return {
          clicked: false,
          courses: [{
            'Name': 'Dummy',
            'Teams': [{
                'Name': 'DummyTeam',
                'Members': [{
                  'Name': 'dummyMember',
                }],
                'Repositories': [{
                  'Name': '',
                  'URL': '',
                  'Issues': [{
                    'Name': '',
                    'Titel': '',
                  }],
                }],
              }],
          }],

          plugins: [{
            'PluginName': 'DummyPlugin1',
            'PluginContent': 'DummyContent'
          },{
            'PluginName': 'DummyPlugin2',
            'PluginContent': 'DummyContent'
          }],
          message: '',
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
        }
      },
      mounted() {
        this.$store.dispatch('user/fetchData', {
          option: "general",
          user: this.getUser,
        }).then(
            (onSuccess) => {
              if (onSuccess.data.success) {
                if (onSuccess.data.courses) {
                  for (const [key, value] of Object.entries(onSuccess.data.courses)) {
                    console.log(`${key}: ${value}`);
                    this.courses = value;
                  }
                }
              }
            },
            (onError) => {
              this.message = onError.toString() || onError.message;
            }
        );
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